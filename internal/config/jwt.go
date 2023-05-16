package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type JwtCfg struct {
	Secret      string `figure:"secret"`
	RefreshLife string `figure:"refresh_life"`
	AccessLife  string `figure:"access_life"`
}

func (c *config) JwtParams() *JwtCfg {
	return c.jwtCfg.Do(func() interface{} {
		var config JwtCfg
		err := figure.
			Out(&config).
			From(kv.MustGetStringMap(c.getter, "jwt")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out jwt params from config"))
		}

		return &config
	}).(*JwtCfg)
}
