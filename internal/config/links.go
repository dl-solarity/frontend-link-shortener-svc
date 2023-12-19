package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Links struct {
	Duration time.Duration `fig:"duration"`
}

func (c *config) Links() Links {
	c.Lock()
	defer c.Unlock()

	if c.links != nil {
		return *c.links
	}

	links := &Links{
		Duration: time.Hour * 24,
	}
	config := kv.MustGetStringMap(c.getter, "links")
	if err := figure.Out(links).From(config).Please(); err != nil {
		panic(errors.Wrap(err, "failed to figure out links"))
	}

	c.links = links

	return *links
}
