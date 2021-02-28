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
	log        = qbsllm.New(qbsllm.Lnormal, "qbsqllm", nil, nil)
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

func (e *entry) addForm(tx *sql.Tx) (id int64) {
	err := tx.QueryRow(`SELECT id FROM form WHERE hash=?`, e.FormHash).Scan(&id)
	switch {
	case err == nil:
		return id
	case err != sql.ErrNoRows:
		log.Fatale(err)
	}
	res, err := tx.Exec(`INSERT INTO form (hash, text) VALUES (?, ?)`,
		e.FormHash,
		e.Form)
	if err != nil {
		log.Fatale(err)
	}
	if id, err = res.LastInsertId(); err != nil {
		log.Fatale(err)
	}
	return id
}

func (e *entry) addEntry(tx *sql.Tx, formId int64) (id int64) {
	// TODO prepare and reuse
	res, err := tx.Exec(`INSERT INTO entry (ts, level, log, code, form)
	                     VALUES (?, ?, ?, ?, ?)`,
		e.TS,
		e.Level,
		e.Log,
		sql.NullString{String: e.Code, Valid: e.Code != ""},
		formId,
	)
	if err != nil {
		log.Fatale(err)
	}
	id, err = res.LastInsertId()
	if err != nil {
		log.Fatale(err)
	}
	for key, vals := range e.Args {
		for _, val := range vals {
			_, err := tx.Exec(`INSERT INTO args (entry, name, value)
			                     VALUES (?, ?, ?)`,
				id,
				key,
				val,
			)
			if err != nil {
				log.Fatale(err)
			}
		}
	}
	return id
}

func readLog(db *sql.DB, rd io.Reader) {
	var fmt bytes.Buffer
	scn := bufio.NewScanner(rd)
	tx, err := db.Begin()
	if err != nil {
		log.Fatale(err)
	}
	defer tx.Commit()
	for scn.Scan() {
		match := lineRegexp.FindStringSubmatch(scn.Text())
		if match == nil {
			log.Fatala("cannot parse `log line`", scn.Text())
			continue
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
		fmt.Reset()
		if e.Args, err = sllm.ParseMap(e.Form, &fmt); err != nil {
			log.Fatale(err)
		}
		e.Form = fmt.String()
		fmtHasher.Reset()
		fmtHasher.Write(fmt.Bytes())
		e.FormHash = fmtHasher.Sum(nil)
		formId := e.addForm(tx)
		e.addEntry(tx, formId)
	}
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
	readLog(db, rd)
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
