package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"database/sql"
	_ "embed"
	"flag"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"git.fractalqb.de/fractalqb/qbsllm"
	"git.fractalqb.de/fractalqb/sllm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	log = qbsllm.New(qbsllm.Lnormal, "qbsqllm", nil, nil)
	// TODO be more precise about source loc to avoid mismatch
	lineRegexp = regexp.MustCompile(`(.{19}) ([^ ]+)[ \t]+\[(.+)\]( \([^)]*\))? (.*)`)
	fmtHasher  = md5.New() // It's supposed to not be security critical, here

	//go:embed sqlite3.sql
	dbCreate string
)

func createDB(db *sql.DB) {
	stmt := dbCreate
	for len(stmt) > 0 {
		sep := strings.Index(stmt, ";\n")
		if sep < 0 {
			break
		}
		if _, err := db.Exec(stmt[:sep+1]); err != nil {
			log.Warne(err)
		}
		stmt = stmt[sep+2:]
	}
}

type entry struct {
	TS       time.Time
	Level    string
	Log      string
	Form     string
	FormHash []byte
	Code     string
	Args     map[string][]string
}

type batch struct {
	tx              *sql.Tx
	stmtSelectForm  *sql.Stmt
	stmtInsertForm  *sql.Stmt
	stmtInsertEntry *sql.Stmt
	stmtInsertArgs  *sql.Stmt
}

func asBatch(db *sql.DB, do func(b *batch) error) (err error) {
	b := new(batch)
	if b.tx, err = db.Begin(); err != nil {
		return err
	}
	defer b.tx.Commit()
	if b.stmtSelectForm, err = b.tx.Prepare(sqlSelectForm); err != nil {
		return err
	}
	defer b.stmtSelectForm.Close()
	if b.stmtInsertForm, err = b.tx.Prepare(sqlInsertForm); err != nil {
		return err
	}
	defer b.stmtInsertForm.Close()
	if b.stmtInsertEntry, err = b.tx.Prepare(sqlInsertEntry); err != nil {
		return err
	}
	defer b.stmtInsertEntry.Close()
	if b.stmtInsertArgs, err = b.tx.Prepare(sqlInsertArgs); err != nil {
		return err
	}
	defer b.stmtInsertArgs.Close()
	err = do(b)
	return err
}

const (
	sqlSelectForm = `SELECT id FROM form WHERE hash=?`
	sqlInsertForm = `INSERT INTO form (hash, text) VALUES (?, ?)`
)

func (e *entry) addForm(b *batch) (id int64, err error) {
	err = b.stmtSelectForm.QueryRow(e.FormHash).Scan(&id)
	switch {
	case err == nil:
		return id, nil
	case err != sql.ErrNoRows:
		return 0, err
	}
	res, err := b.stmtInsertForm.Exec(e.FormHash, e.Form)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

const (
	sqlInsertEntry = `INSERT INTO entry (ts, level, log, code, form, file, line)
	                  VALUES (?, ?, ?, ?, ?, ?, ?)`
	sqlInsertArgs = `INSERT INTO args (entry, name, value)
			         VALUES (?, ?, ?)`
)

func (e *entry) addEntry(b *batch, formId int64, file, line int32) (id int64, err error) {
	res, err := b.stmtInsertEntry.Exec(
		e.TS,
		e.Level,
		e.Log,
		sql.NullString{String: e.Code, Valid: e.Code != ""},
		formId,
		sql.NullInt32{Int32: file, Valid: file > 0},
		line,
	)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for key, vals := range e.Args {
		for _, val := range vals {
			_, err := b.stmtInsertArgs.Exec(
				id,
				key,
				val,
			)
			if err != nil {
				log.Fatale(err)
			}
		}
	}
	return id, nil
}

func readLog(b *batch, rd io.Reader, fileId int32) error {
	var form bytes.Buffer
	scn := bufio.NewScanner(rd)
	line := int32(0)
	for scn.Scan() {
		line++
		match := lineRegexp.FindStringSubmatch(scn.Text())
		if match == nil {
			return sllm.Error("cannot parse `log line`", scn.Text())
		}
		ts, err := time.Parse("Jan _2 15:04:05.000", match[1])
		if err != nil {
			log.Fatale(err)
		}
		e := entry{
			TS:    ts,
			Level: match[2],
			Log:   match[3],
			Code:  codeEntry(match[4]),
			Form:  match[5],
		}
		form.Reset()
		if e.Args, err = sllm.ParseMap(e.Form, &form); err != nil {
			return sllm.Error("sllm `parse error`", err)
		}
		e.Form = form.String()
		fmtHasher.Reset()
		fmtHasher.Write(form.Bytes())
		e.FormHash = fmtHasher.Sum(nil)
		formId, err := e.addForm(b)
		e.addEntry(b, formId, fileId, line)
	}
	return nil
}

func codeEntry(raw string) string {
	return strings.Trim(raw, "() ")
}

func readLogFile(db *sql.DB, file string) {
	rd, err := os.Open(file)
	if err != nil {
		log.Fatale(err)
	}
	defer rd.Close()
	err = asBatch(db, func(b *batch) error {
		res, err := b.tx.Exec(`INSERT INTO file (name) VALUES (?)`, file)
		if err != nil {
			return err
		}
		fileId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		return readLog(b, rd, int32(fileId))
	})
	if err != nil {
		log.Fatale(err)
	}
}

func main() {
	dbfile := flag.String("db", "qbsqllm.db", "Set name of sqlite3 file")
	flag.Parse()
	db, err := sql.Open("sqlite3", *dbfile)
	if err != nil {
		log.Fatale(err)
	}
	defer db.Close()
	createDB(db)
	for _, arg := range flag.Args() {
		readLogFile(db, arg)
	}
}
