package lib

import (
	"errors"

	l "git.fractalqb.de/fractalqb/qbsllm"
)

var (
	log    = l.New(l.Lnormal, "qbsllm-example-lib", nil, nil)
	LogCfg = l.Config(log)
)

func Yalf(s string, f bool) error {
	if len(s) == 0 {
		err := errors.New("passed empty string to yalf")
		log.Error(l.Err(err))
		return err
	}
	log.DebugA("yalf called with `s`and `f`", s, f)
	return nil
}
