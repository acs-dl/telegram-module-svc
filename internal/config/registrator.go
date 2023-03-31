package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RegistratorConfig struct {
	OuterUrl string `fig:"outer_url,required"`
	InnerUrl string `fig:"inner_url,required"`
	Title    string `fig:"title,required"`
	Topic    string `fig:"topic,required"`
	Prefix   string `fig:"prefix,required"`
}

func (c *config) Registrator() RegistratorConfig {
	return c.registrator.Do(func() interface{} {
		cfg := RegistratorConfig{}

		if err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "registrator")).
			Please(); err != nil {
			panic(errors.Wrap(err, "failed to get core registrator config from config"))
		}

		return cfg
	}).(RegistratorConfig)
}
