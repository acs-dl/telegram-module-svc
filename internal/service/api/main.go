package api

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/config"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type apiRouter struct {
	cfg config.Config
}

func (r *apiRouter) run() error {
	router := r.apiRouter()

	if err := r.cfg.Copus().RegisterChi(router); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(r.cfg.Listener(), router)
}

func NewApiRouter(cfg config.Config) *apiRouter {
	return &apiRouter{cfg: cfg}
}

func Run(_ context.Context, cfg config.Config) {
	if err := NewApiRouter(cfg).run(); err != nil {
		panic(err)
	}
}
