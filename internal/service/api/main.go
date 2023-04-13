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

func (r *Router) run() error {
	router := r.apiRouter()

	if err := r.cfg.Copus().RegisterChi(router); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(r.cfg.Listener(), router)
}

func NewApiRouter(ctx context.Context, cfg config.Config) *Router {
	return &Router{ctx: ctx, cfg: cfg}
}

func Run(ctx context.Context, cfg config.Config) {
	if err := NewApiRouter(ctx, cfg).run(); err != nil {
		panic(err)
	}
}
