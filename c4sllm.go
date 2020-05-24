package qbsllm

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/c4qblog"
)

func NewConfig(log *Logger) c4hgol.Configurer {
	return c4qblog.NewConfig(&log.Logger)
}

// func Config(log *Logger, sub ...c4hgol.Configurer) c4hgol.ShowSrcConfigurer {
// 	cfg := c4qblog.NewConfig(&log.Logger)
// 	if len(sub) == 0 {
// 		return cfg
// 	}
// 	topic := c4hgol.NewTopic(cfg, sub...)
// 	return topic
// }
