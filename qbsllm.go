package qbsllm

import (
	"context"
	"io"

	"git.fractalqb.de/fractalqb/qblog"
	"git.fractalqb.de/fractalqb/sllm"
)

type Logger qblog.Logger

func New(level qblog.Level, title string, wr io.Writer, fmt qblog.Formatter) *Logger {
	res := qblog.New(level, title, wr, fmt)
	return (*Logger)(res)
}

const calldepth = 3

func (l *Logger) CtxWr(level qblog.Level, ctx context.Context, message string) {
	(*qblog.Logger)(l).Out(calldepth, level, ctx, qblog.Msg(message))
}

func (l *Logger) Wr(level qblog.Level, message string) {
	(*qblog.Logger)(l).Out(calldepth, level, nil, qblog.Msg(message))
}

func (l *Logger) CtxArgs(level qblog.Level, ctx context.Context, tmpl string, args ...interface{}) {
	(*qblog.Logger)(l).Out(calldepth, level, ctx, sllm.Args(tmpl, args...))
}

func (l *Logger) Args(level qblog.Level, tmpl string, args ...interface{}) {
	(*qblog.Logger)(l).Out(calldepth, level, nil, sllm.Args(tmpl, args...))
}

func (l *Logger) CtxMap(level qblog.Level, ctx context.Context, tmpl string, args sllm.ArgMap) {
	(*qblog.Logger)(l).Out(calldepth, level, ctx, sllm.Map(tmpl, args))
}

func (l *Logger) Map(level qblog.Level, tmpl string, args sllm.ArgMap) {
	(*qblog.Logger)(l).Out(calldepth, level, nil, sllm.Map(tmpl, args))
}
