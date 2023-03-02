package api

import (
	"fmt"
	auth "gitlab.com/distributed_lab/acs/auth/middlewares"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data/postgres"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (r *apiRouter) apiRouter() chi.Router {
	router := chi.NewRouter()

	logger := r.cfg.Log().WithField("service", fmt.Sprintf("%s-api", data.ModuleName))

	secret := r.cfg.JwtParams().Secret

	router.Use(
		ape.RecoverMiddleware(logger),
		ape.LoganMiddleware(logger),
		ape.CtxMiddleware(
			//base
			handlers.CtxLog(logger),

			// storage
			handlers.CtxPermissionsQ(postgres.NewPermissionsQ(r.cfg.DB())),
			handlers.CtxUsersQ(postgres.NewUsersQ(r.cfg.DB())),
			handlers.CtxLinksQ(postgres.NewLinksQ(r.cfg.DB())),

			// connectors

			// other configs
			handlers.CtxParams(r.cfg.Telegram()),
		),
	)

	router.Route("/integrations/telegram", func(r chi.Router) {
		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Member]}...)).
			Get("/get_input", handlers.GetInputs)
		//r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Member]}...)).
		r.Get("/get_available_roles", handlers.GetRoles)

		r.Get("/role", handlers.GetRole)      // comes from orchestrator
		r.Get("/roles", handlers.GetRolesMap) // comes from orchestrator

		r.Route("/links", func(r chi.Router) {
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin]}...)).
				Post("/", handlers.AddLink)
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin]}...)).
				Delete("/", handlers.RemoveLink)
		})

		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Member]}...)).
			Get("/permissions", handlers.GetPermissions)

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", handlers.GetUserById) // comes from orchestrator

			//r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Member]}...)).
			r.Get("/", handlers.GetUsers)
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Member]}...)).
				Get("/unverified", handlers.GetUnverifiedUsers)
		})
	})

	return router
}
