package api

import (
	"fmt"

	auth "github.com/acs-dl/auth-svc/middlewares"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/data/postgres"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (r *Router) apiRouter() chi.Router {
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
			handlers.CtxChatsQ(postgres.NewChatsQ(r.cfg.DB())),
			handlers.CtxLinksQ(postgres.NewLinksQ(r.cfg.DB())),

			// connectors

			// other configs
			handlers.CtxParentContext(r.ctx),
		),
	)

	router.Route("/integrations/telegram", func(r chi.Router) {
		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
			Get("/get_input", handlers.GetInputs)
		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
			Get("/get_available_roles", handlers.GetRoles)

		r.Get("/role", handlers.GetRole)               // comes from orchestrator
		r.Get("/roles", handlers.GetRolesMap)          // comes from orchestrator
		r.Get("/user_roles", handlers.GetUserRolesMap) // comes from orchestrator

		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
			Get("/submodule", handlers.CheckSubmodule)

		r.Route("/links", func(r chi.Router) {
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner]}...)).
				Post("/", handlers.AddLink)
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner]}...)).
				Delete("/", handlers.RemoveLink)
		})

		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
			Route("/estimate_refresh", func(r chi.Router) {
				r.Post("/submodule", handlers.GetEstimatedRefreshSubmodule)
				r.Post("/module", handlers.GetEstimatedRefreshModule)
			})

		r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
			Get("/permissions", handlers.GetPermissions)

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", handlers.GetUserById) // comes from orchestrator

			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
				Get("/", handlers.GetUsers)
			r.With(auth.Jwt(secret, data.ModuleName, []string{data.Roles[data.Admin], data.Roles[data.Owner], data.Roles[data.Member]}...)).
				Get("/unverified", handlers.GetUnverifiedUsers)
		})
	})

	return router
}
