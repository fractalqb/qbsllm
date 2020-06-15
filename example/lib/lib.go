package lib

import (
	"errors"

	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "qbsllm-example-lib", nil, nil)
	LogCfg = c4hgol.Config(qbsllm.NewConfig(log))
)

func Yalf(s string, f bool) error {
	if len(s) == 0 {
		err := errors.New("passed empty string to yalf")
		log.Errore(err)
		return err
	}
	log.Debuga("yalf called with `s` and `f`", s, f)
	return nil
}
