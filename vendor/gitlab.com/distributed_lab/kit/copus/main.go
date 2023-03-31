package copus

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus/cop"
	"gitlab.com/distributed_lab/kit/copus/janus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// List of Copuser implementations
var serviceSlice = map[string]func(getter kv.Getter) types.Copuser{
	"cop":   cop.NewCoper,
	"janus": janus.NewJanuserWrapper,
}

// NewCopuser returns new Copuser instance
func NewCopuser(getter kv.Getter) types.Copuser {
	return &generalCopuser{
		getter: getter,
	}
}

type generalCopuser struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *generalCopuser) Copus() types.Copus {
	return c.once.Do(func() interface{} {
		var copuser types.Copuser

		for name, comfig := range serviceSlice {
			if hasKey(c.getter, name) {
				if copuser == nil {
					copuser = comfig(c.getter)
				} else {
					panic(errors.Wrap(
						errors.New("failed to create copuser"),
						"conflicting keys specified, check that only one of specified keys presented in your config",
						logan.F{"keys": keys(serviceSlice)},
					))
				}
			}
		}

		if copuser == nil {
			panic(errors.Wrap(
				errors.New("failed to create copuser"),
				"failed to get any of keys",
				logan.F{"keys": keys(serviceSlice)},
			))
		}

		return copuser.Copus()
	}).(types.Copus)
}

func keys(source map[string]func(getter kv.Getter) types.Copuser) []string {
	keys := make([]string, 0, len(source))
	for key := range serviceSlice {
		keys = append(keys, key)
	}
	return keys
}

func hasKey(getter kv.Getter, key string) bool {
	raw, err := getter.GetStringMap(key)
	return err == nil && raw != nil && len(raw) != 0
}
