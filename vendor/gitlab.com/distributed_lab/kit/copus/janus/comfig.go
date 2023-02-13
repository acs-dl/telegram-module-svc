package janus

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/janus"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
)

func NewJanuserWrapper(getter kv.Getter) types.Copuser {
	return &januserWrapper{
		janus: janus.NewJanuser(getter),
	}
}

type januserWrapper struct {
	janus janus.Januser
}

func (j *januserWrapper) Copus() types.Copus {
	return &janusWrapper{
		janus: j.janus.Janus(),
	}
}

type janusWrapper struct {
	janus *janus.Janus
}

func (j *janusWrapper) RegisterGojiEndpoint(endpoint, method string) error {
	return j.janus.RegisterGojiEndpoint(endpoint, method)
}

func (j *janusWrapper) RegisterChi(r chi.Router) error {
	return j.janus.RegisterChi(r)
}

func (j *janusWrapper) WithLog(log *logan.Entry) types.Copus {
	j.janus = j.janus.WithLog(log)
	return j
}
