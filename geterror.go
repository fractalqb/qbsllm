package qbsllm

import (
	"errors"
	"strings"

	"git.fractalqb.de/fractalqb/qblog"
	"git.fractalqb.de/fractalqb/sllm"
)

type errLogger Logger

// Get returns a variant of Logger that only supports the Error* methods to log
// errors. The new variants return the message part of the log entry as an
// error object.
func (l *Logger) Get() *errLogger { return (*errLogger)(l) }

func (l *errLogger) Errors(msg string) error {
	l.Out(calldepth, qblog.Lerror, qblog.Str(msg))
	return errors.New(msg)
}

func (l *errLogger) Errore(err error) error {
	l.Out(calldepth, qblog.Lerror, qblog.Err(err))
	return err
}

func (l *errLogger) Errorf(format string, a ...interface{}) error {
	wr := qblog.Fmt(format, a...)
	var sb strings.Builder
	wr.WriteTo(&sb)
	l.Out(calldepth, qblog.Lerror, qblog.Str(sb.String()))
	return errors.New(sb.String())
}

func (l *errLogger) Errora(tmpl string, args ...interface{}) error {
	wr := wrSllm{(*Logger)(l), tmpl, sllm.Args(nil, args...)}
	var sb strings.Builder
	wr.WriteTo(&sb)
	l.Out(calldepth, qblog.Lerror, qblog.Str(sb.String()))
	return errors.New(sb.String())
}

func (l *errLogger) Errorm(tmpl string, args sllm.ArgMap) error {
	wr := wrSllm{(*Logger)(l), tmpl, sllm.Map(nil, args)}
	var sb strings.Builder
	wr.WriteTo(&sb)
	l.Out(calldepth, qblog.Lerror, qblog.Str(sb.String()))
	return errors.New(sb.String())
}
