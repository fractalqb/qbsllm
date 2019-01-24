package qbsllm

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/c4qblog"
)

func Config(log *Logger, sub ...c4hgol.Configurer) c4hgol.Configurer {
	cfg := c4qblog.Config(&log.Logger)
	if len(sub) == 0 {
		return cfg
	}
	topic := c4hgol.Topic{
		TopicName:  cfg.Name(),
		Configures: append(sub, cfg),
	}
	return topic
}
