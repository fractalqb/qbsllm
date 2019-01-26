package qbsllm

import (
	"io"

	"git.fractalqb.de/fractalqb/qblog"
	"git.fractalqb.de/fractalqb/sllm"
)

const (
	FsrcLoc  = qblog.FsrcLoc
	FsrcPath = qblog.FsrcPath
)

type Level = qblog.Level

const (
	LmostIrrelevant Level = qblog.LmostIrrelevant
	Lnormal         Level = qblog.Lnormal
	LmostImportant  Level = qblog.LmostImportant
	Loff            Level = qblog.Loff

	Lpanic Level = qblog.Lpanic
	Lfatal Level = qblog.Lfatal

	Ltrace Level = qblog.Ltrace
	Ldebug Level = qblog.Ldebug
	Linfo  Level = qblog.Linfo
	Lwarn  Level = qblog.Lwarn
	Lerror Level = qblog.Lerror
)

type Formatter = qblog.Formatter
type Str = qblog.Str

func Fmt(format string, a ...interface{}) qblog.FmtMsg { return qblog.Fmt(format, a...) }

func Err(err error) qblog.ErrMsg { return qblog.Err(err) }

type Logger struct {
	qblog.Logger
}

func New(level Level, title string, wr io.Writer, fmt Formatter) *Logger {
	tmp := qblog.New(level, title, wr, fmt)
	res := &Logger{*tmp}
	return res
}

const calldepth = 3

func (l *Logger) Args(level qblog.Level, tmpl string, args ...interface{}) {
	l.Out(calldepth, level, sllm.Args(tmpl, args...))
}

func (l *Logger) Map(level qblog.Level, tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, level, sllm.Map(tmpl, args))
}

func (l *Logger) TraceA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ltrace, sllm.Args(tmpl, args...))
}

func (l *Logger) TraceM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ltrace, sllm.Map(tmpl, args))
}

func (l *Logger) DebugA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ldebug, sllm.Args(tmpl, args...))
}

func (l *Logger) DebugM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ldebug, sllm.Map(tmpl, args))
}

func (l *Logger) InfoA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Linfo, sllm.Args(tmpl, args...))
}

func (l *Logger) InfoM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Linfo, sllm.Map(tmpl, args))
}

func (l *Logger) WarnA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lwarn, sllm.Args(tmpl, args...))
}

func (l *Logger) WarnM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lwarn, sllm.Map(tmpl, args))
}

func (l *Logger) ErrorA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lerror, sllm.Args(tmpl, args...))
}

func (l *Logger) ErrorM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lerror, sllm.Map(tmpl, args))
}

func (l *Logger) PanicA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lpanic, sllm.Args(tmpl, args...))
}

func (l *Logger) PanicM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lpanic, sllm.Map(tmpl, args))
}

func (l *Logger) Fatal(message io.WriterTo) {
	l.Out(calldepth, qblog.Lfatal, message)
}

func (l *Logger) FatalA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lfatal, sllm.Args(tmpl, args...))
}
