package main

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"git.fractalqb.de/fractalqb/qbsllm/example/lib"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "qbsllm-example-exe", nil, nil)
	logCfg = c4hgol.Config(qbsllm.NewConfig(log), lib.LogCfg)
)

func main() {
	logCfg.SetLevel(c4hgol.LeastImportant)
	log.SetFlags(qbsllm.FsrcLoc)
	log.Infos("gonna call lib.Yalf twice")
	lib.Yalf("", false)
	lib.Yalf("second time", true)
	log.Infos("DONE calling lib.Yalf")
}
