package janus

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Januser interface {
	Janus() *Janus
}

// Deprecated: use copus.NewCopuser instead of it
func NewJanuser(getter kv.Getter) Januser {
	return &januser{
		getter: getter,
	}
}

type januser struct {
	getter kv.Getter
	once   comfig.Once
}

func (j *januser) Janus() *Janus {
	return j.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(j.getter, "janus")

		var probe struct {
			Disabled bool `fig:"disabled"`
		}

		if err := figure.Out(&probe).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out janus probe"))
		}

		if probe.Disabled {
			return NewNoOp()
		}

		var config struct {
			Endpoint string `fig:"endpoint,required"`
			Upstream string `fig:"upstream,required"`
			Surname  string `fig:"surname,required"`
		}

		if err := figure.Out(&config).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out janus"))
		}

		return New(config.Endpoint, Upstream{
			Target:  config.Upstream,
			Surname: config.Surname,
		})
	}).(*Janus)
}
