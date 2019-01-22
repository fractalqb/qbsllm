package qbsllm

import (
	"io"

	"git.fractalqb.de/fractalqb/qblog"
	"git.fractalqb.de/fractalqb/sllm"
)

type Logger struct {
	qblog.Logger
}

func New(level qblog.Level, title string, wr io.Writer, fmt qblog.Formatter) *Logger {
	tmp := qblog.New(level, title, wr, fmt)
	res := &Logger{*tmp}
	return res
}

const calldepth = 3

func (l *Logger) Msg(level qblog.Level, message string) {
	l.Out(calldepth, level, qblog.Msg(message))
}

func (l *Logger) Args(level qblog.Level, tmpl string, args ...interface{}) {
	l.Out(calldepth, level, sllm.Args(tmpl, args...))
}

func (l *Logger) Map(level qblog.Level, tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, level, sllm.Map(tmpl, args))
}

func (l *Logger) Trace(message string) {
	l.Out(calldepth, qblog.Ltrace, qblog.Msg(message))
}

func (l *Logger) TraceA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ltrace, sllm.Args(tmpl, args...))
}

func (l *Logger) TraceM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ltrace, sllm.Map(tmpl, args))
}

func (l *Logger) Debug(message string) {
	l.Out(calldepth, qblog.Ldebug, qblog.Msg(message))
}

func (l *Logger) DebugA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ldebug, sllm.Args(tmpl, args...))
}

func (l *Logger) DebugM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ldebug, sllm.Map(tmpl, args))
}

func (l *Logger) Info(message string) {
	l.Out(calldepth, qblog.Linfo, qblog.Msg(message))
}

func (l *Logger) InfoA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Linfo, sllm.Args(tmpl, args...))
}

func (l *Logger) InfoM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Linfo, sllm.Map(tmpl, args))
}

func (l *Logger) Warn(message string) {
	l.Out(calldepth, qblog.Lwarn, qblog.Msg(message))
}

func (l *Logger) WarnA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lwarn, sllm.Args(tmpl, args...))
}

func (l *Logger) WarnM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lwarn, sllm.Map(tmpl, args))
}

func (l *Logger) Error(message string) {
	l.Out(calldepth, qblog.Lerror, qblog.Msg(message))
}

func (l *Logger) ErrorA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lerror, sllm.Args(tmpl, args...))
}

func (l *Logger) ErrorM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lerror, sllm.Map(tmpl, args))
}

func (l *Logger) Panic(message string) {
	l.Out(calldepth, qblog.Lpanic, qblog.Msg(message))
}

func (l *Logger) PanicA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lpanic, sllm.Args(tmpl, args...))
}

func (l *Logger) PanicM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lpanic, sllm.Map(tmpl, args))
}

func (l *Logger) Fatal(message string) {
	l.Out(calldepth, qblog.Lfatal, qblog.Msg(message))
}

func (l *Logger) FatalA(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lfatal, sllm.Args(tmpl, args...))
}

func (l *Logger) FatalM(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lfatal, sllm.Map(tmpl, args))
}
