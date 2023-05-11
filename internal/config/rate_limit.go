package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RateLimitCfg struct {
	RequestsAmount int64         `fig:"requests_amount,required"`
	TimeLimit      time.Duration `fig:"time_limit,required"`
}

func (c *config) RateLimit() *RateLimitCfg {
	return c.rateLimit.Do(func() interface{} {
		var cfg RateLimitCfg
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "rate_limit")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out rate limit params from config"))
		}

		return &cfg
	}).(*RateLimitCfg)
}
