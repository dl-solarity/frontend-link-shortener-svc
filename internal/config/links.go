package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LinksConfig struct {
	Duration time.Duration `fig:"duration"`
}

type Links interface {
	LinksConfig() *LinksConfig
}

func NewLinks(getter kv.Getter) Links {
	return &links{
		getter: getter,
	}
}

type links struct {
	getter kv.Getter
	once   comfig.Once
}

func (l *links) LinksConfig() *LinksConfig {
	return l.once.Do(func() interface{} {
		config := &LinksConfig{
			Duration: time.Hour * 24,
		}
		raw := kv.MustGetStringMap(l.getter, "links")
		if err := figure.Out(config).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out links"))
		}

		return config
	}).(*LinksConfig)
}
