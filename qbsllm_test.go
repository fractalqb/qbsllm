package qbsllm

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"git.fractalqb.de/fractalqb/qblog"
)

func testFormat(
	buf *bytes.Buffer,
	now time.Time,
	lvl qblog.Level,
	title, file string,
	line int,
	ctx qblog.Context,
	message io.WriterTo,
) {
	fmt.Fprintf(buf, "%s %s %s:%d <%s>",
		title,
		qblog.LevelName(lvl),
		file, line,
		message)
}

func ExampleNewApi() {
	l := New(qblog.Linfo, "qbsllm", os.Stdout, testFormat)
	l.SetFlags(qblog.FsrcLoc)
	l.InfoA("just a `adj` message", "fancy")
	// Output:
	// qbsllm INFO qbsllm_test.go:32 <just a `adj:fancy` message>
}
