package main

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	l "git.fractalqb.de/fractalqb/qbsllm"
	"git.fractalqb.de/fractalqb/qbsllm/example/lib"
)

var (
	log    = l.New(l.Lnormal, "qbsllm-example-exe", nil, nil)
	logCfg = l.Config(log, lib.LogCfg)
)

func main() {
	logCfg.SetLevel(c4hgol.LeastImportant)
	log.SetFlags(l.FsrcLoc)
	log.Info(l.Msg("gonna call lib.Yalf twice"))
	lib.Yalf("", false)
	lib.Yalf("second time", true)
	log.Info(l.Msg("DONE calling lib.Yalf"))
}
