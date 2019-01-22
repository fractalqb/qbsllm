package qbsllm

import (
	"os"

	"git.fractalqb.de/fractalqb/qblog"
)

func ExampleOldNewApi() {
	l := New(qblog.Linfo, "qbsllm", os.Stdout, qblog.BasicFormat)
	l.Args(qblog.Linfo, "just a `adj` message", "fancy")
	// Output:
}
