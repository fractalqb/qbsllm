package log

import (
	"os"
	"path/filepath"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
)

var (
	root = qbsllm.New(qbsllm.Lnormal, title(), nil, nil)
	Cfg  = c4hgol.Config(qbsllm.NewConfig(root))
)

func title() string {
	res := filepath.Base(os.Args[0])
	if ext := filepath.Ext(res); len(ext) > 0 {
		res = res[:len(res)-len(ext)]
	}
	return res
}

func Debugs(msg string)                       { root.Debugs(msg) }
func Debugf(fmt string, args ...interface{})  { root.Debugf(fmt, args...) }
func Debuga(tmpl string, args ...interface{}) { root.Debuga(tmpl, args...) }

func Infos(msg string)                       { root.Infos(msg) }
func Infof(fmt string, args ...interface{})  { root.Infof(fmt, args...) }
func Infoa(tmpl string, args ...interface{}) { root.Infoa(tmpl, args...) }

func Warns(msg string)                       { root.Warns(msg) }
func Warnf(fmt string, args ...interface{})  { root.Warnf(fmt, args...) }
func Warna(tmpl string, args ...interface{}) { root.Warna(tmpl, args...) }

func Errors(msg string)                       { root.Errors(msg) }
func Errore(err error)                        { root.Errore(err) }
func Errora(tmpl string, args ...interface{}) { root.Errora(tmpl, args...) }

func Panica(tmpl string, args ...interface{}) { root.Panica(tmpl, args...) }
func Panice(err error)                        { root.Panice(err) }
func Panics(msg string)                       { root.Panics(msg) }

func Fatala(tmpl string, args ...interface{}) { root.Fatala(tmpl, args...) }
func Fatale(err error)                        { root.Fatale(err) }
func Fatals(msg string)                       { root.Fatals(msg) }
