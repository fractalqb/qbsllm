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

func Infos(msg string) { root.Infos(msg) }

func Infoa(tmpl string, args ...interface{}) { root.Infoa(tmpl, args...) }

func Errora(tmpl string, args ...interface{}) { root.Errora(tmpl, args...) }

func Fatala(tmpl string, args ...interface{}) { root.Fatala(tmpl, args...) }

func Fatale(err error) { root.Fatale(err) }

func Fatals(msg string) { root.Fatals(msg) }
