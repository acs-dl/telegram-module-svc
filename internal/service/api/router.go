package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (r *apiRouter) apiRouter() chi.Router {
	router := chi.NewRouter()

	logger := r.cfg.Log().WithField("service", fmt.Sprintf("%s-api", data.ModuleName))

	router.Use(
		ape.RecoverMiddleware(logger),
		ape.LoganMiddleware(logger),
		ape.CtxMiddleware(
			//base
			handlers.CtxLog(logger),

			// storage

			// connectors

			// other configs
		),
	)

	router.Route("/integrations/telegram", func(r chi.Router) {

	})

	return router
}
