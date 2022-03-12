package qbsllm

import (
	"io"

	"git.fractalqb.de/fractalqb/qblog"
	"git.fractalqb.de/fractalqb/sllm"
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
	SllmDeco func(sllm.ParamWriter) sllm.ParamWriter
}

func New(level Level, title string, wr io.Writer, fmt Formatter) *Logger {
	tmp := qblog.New(level, title, wr, fmt)
	res := &Logger{*tmp, nil}
	return res
}

const calldepth = 3

func (l *Logger) Str(level Level, msg string) {
	l.Out(calldepth, level, qblog.Str(msg))
}

func (l *Logger) Err(level Level, err error) error {
	l.Out(calldepth, level, qblog.Err(err))
	return err
}

func (l *Logger) Fmt(level Level, format string, a ...interface{}) {
	l.Out(calldepth, level, qblog.Fmt(format, a))
}

func (l *Logger) Args(level Level, tmpl string, args ...interface{}) {
	l.Out(calldepth, level, wrSllm{
		log:  l,
		tmpl: tmpl,
		swp:  sllm.Args(nil, args...),
	})
}

func (l *Logger) Map(level Level, tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, level, wrSllm{
		log:  l,
		tmpl: tmpl,
		swp:  sllm.Map(nil, args),
	})
}

func (l *Logger) Traces(msg string) {
	l.Out(calldepth, qblog.Ltrace, qblog.Str(msg))
}

func (l *Logger) Tracee(err error) {
	l.Out(calldepth, qblog.Ltrace, qblog.Err(err))
}

func (l *Logger) Tracef(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Ltrace, qblog.Fmt(format, a...))
}

func (l *Logger) Tracea(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ltrace, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Tracem(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ltrace, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Debugs(msg string) {
	l.Out(calldepth, qblog.Ldebug, qblog.Str(msg))
}

func (l *Logger) Debuge(err error) {
	l.Out(calldepth, qblog.Ldebug, qblog.Err(err))
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Ldebug, qblog.Fmt(format, a...))
}

func (l *Logger) Debuga(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Ldebug, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Debugm(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Ldebug, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Infos(msg string) {
	l.Out(calldepth, qblog.Linfo, qblog.Str(msg))
}

func (l *Logger) Infoe(err error) {
	l.Out(calldepth, qblog.Linfo, qblog.Err(err))
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Linfo, qblog.Fmt(format, a...))
}

func (l *Logger) Infoa(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Linfo, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Infom(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Linfo, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Warns(msg string) {
	l.Out(calldepth, qblog.Lwarn, qblog.Str(msg))
}

func (l *Logger) Warne(err error) {
	l.Out(calldepth, qblog.Lwarn, qblog.Err(err))
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Lwarn, qblog.Fmt(format, a...))
}

func (l *Logger) Warna(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lwarn, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Warnm(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lwarn, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Errors(msg string) {
	l.Out(calldepth, qblog.Lerror, qblog.Str(msg))
}

func (l *Logger) Errore(err error) {
	l.Out(calldepth, qblog.Lerror, qblog.Err(err))
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Lerror, qblog.Fmt(format, a...))
}

func (l *Logger) Errora(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lerror, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Errorm(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lerror, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Panics(msg string) {
	l.Out(calldepth, qblog.Lpanic, qblog.Str(msg))
}

func (l *Logger) Panice(err error) {
	l.Out(calldepth, qblog.Lpanic, qblog.Err(err))
}

func (l *Logger) Panicf(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Lpanic, qblog.Fmt(format, a...))
}

func (l *Logger) Panica(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lpanic, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Panicm(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lpanic, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

func (l *Logger) Fatals(msg string) {
	l.Out(calldepth, qblog.Lfatal, qblog.Str(msg))
}

func (l *Logger) Fatale(err error) {
	l.Out(calldepth, qblog.Lfatal, qblog.Err(err))
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Out(calldepth, qblog.Lfatal, qblog.Fmt(format, a...))
}

func (l *Logger) Fatala(tmpl string, args ...interface{}) {
	l.Out(calldepth, qblog.Lfatal, wrSllm{l, tmpl, sllm.Args(nil, args...)})
}

func (l *Logger) Fatalm(tmpl string, args sllm.ArgMap) {
	l.Out(calldepth, qblog.Lpanic, wrSllm{l, tmpl, sllm.Map(nil, args)})
}

type wrSllm struct {
	log  *Logger
	tmpl string
	swp  sllm.ParamWriter
}

func (wst wrSllm) WriteTo(w io.Writer) (n int64, err error) {
	var wn int
	if wst.log.SllmDeco == nil {
		wn, err = sllm.Expand(w, wst.tmpl, wst.swp)
	} else {
		pw := wst.log.SllmDeco(wst.swp)
		wn, err = sllm.Expand(w, wst.tmpl, pw)
	}
	return int64(wn), err
}
