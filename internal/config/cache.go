package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"time"
)

type CacheCfg struct {
	defaultExpiration string `fig:"default_expiration,required"`
	cleanupInterval   string `fig:"cleanup_interval,required"`
}

type CacheData struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

func (c *config) Cache() *CacheData {
	return c.gitlab.Do(func() interface{} {
		var cfg CacheCfg
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "cache")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out cache params from config"))
		}

		return &CacheData{
			DefaultExpiration: convertToDuration(cfg.defaultExpiration),
			CleanupInterval:   convertToDuration(cfg.cleanupInterval),
		}
	}).(*CacheData)
}

func convertToDuration(durationString string) time.Duration {
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		panic(errors.Wrap(err, "failed to convert duration"))
	}

	return duration
}
