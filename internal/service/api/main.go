package api

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Router struct {
	cfg config.Config
	ctx context.Context
}

func (r *Router) Run() error {
	router := r.apiRouter()

	if err := r.cfg.Copus().RegisterChi(router); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(r.cfg.Listener(), router)
}

func NewRouterAsInterface(cfg config.Config, ctx context.Context) interface{} {
	return interface{}(&Router{
		cfg: cfg,
		ctx: ctx,
	})
}

func RunRouterAsInterface(structure interface{}, _ context.Context) {
	err := (structure.(*Router)).Run()
	if err != nil {
		panic(err)
	}
}
