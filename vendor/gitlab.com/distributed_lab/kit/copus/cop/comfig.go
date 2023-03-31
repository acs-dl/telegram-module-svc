package cop

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewCoper(getter kv.Getter) types.Copuser {
	return &coper{
		getter: getter,
	}
}

type coper struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *coper) Copus() types.Copus {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "cop")

		var probe struct {
			Disabled bool `fig:"disabled"`
		}

		if err := figure.Out(&probe).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out cop probe"))
		}

		if probe.Disabled {
			return NewNoOp()
		}

		var config CopConfig

		if err := figure.Out(&config).From(raw).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out cop"))
		}

		return New(config)
	}).(*Cop)
}
