package config

import (
	"time"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RunnersCfg struct {
	Registrar time.Duration `fig:"registrar,required"`
	Worker    time.Duration `fig:"worker,required"`
	Receiver  time.Duration `fig:"receiver,required"`
	Sender    time.Duration `fig:"sender,required"`
}

func (c *config) Runners() *RunnersCfg {
	return c.runners.Do(func() interface{} {
		var cfg RunnersCfg
		err := figure.
			Out(&cfg).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "runners")).
			Please()

		if err != nil {
			panic(errors.Wrap(err, "failed to figure out runners params from config"))
		}

		return &cfg
	}).(*RunnersCfg)
}
