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
	lvl Level,
	title, file string,
	line int,
	ctx qblog.Context,
	message io.WriterTo,
) (needNewline bool) {
	fmt.Fprintf(buf, "%s %s <",
		title,
		qblog.LevelName(lvl))
	message.WriteTo(buf)
	buf.WriteString(">\n")
	return false
}

func ExampleNewApi() {
	l := New(Linfo, "qbsllm", os.Stdout, testFormat)
	l.Infoa("just a `adj` message", "fancy")
	// Output:
	// qbsllm INFO <just a `adj:fancy` message>
}
