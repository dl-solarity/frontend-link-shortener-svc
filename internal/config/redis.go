package config

import (
	"context"

	redis2 "github.com/redis/go-redis/v9"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type redisConfig struct {
	URL string `fig:"url,required"`
}

type Redis interface {
	RedisClient() *redis2.Client
}

func NewRedis(getter kv.Getter) Redis {
	return &redis{
		getter: getter,
	}
}

type redis struct {
	getter kv.Getter
	once   comfig.Once
}

func (r *redis) readConfig() redisConfig {
	config := redisConfig{}

	err := figure.Out(&config).
		From(kv.MustGetStringMap(r.getter, "redis")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out redis"))
	}

	return config
}

func (r *redis) RedisClient() *redis2.Client {
	return r.once.Do(func() interface{} {
		config := r.readConfig()

		opts, err := redis2.ParseURL(config.URL)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse redis url"))
		}

		cli := redis2.NewClient(opts)
		_, err = cli.Ping(context.Background()).Result()
		if err != nil {
			panic(errors.Wrap(err, "failed to ping redis database"))
		}

		return cli
	}).(*redis2.Client)
}
