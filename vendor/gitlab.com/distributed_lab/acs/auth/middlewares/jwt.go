package middleware

import (
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/acs/auth/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Jwt(secret, module string, permissions ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			splitAuthHeader := strings.Split(authHeader, " ")
			if len(splitAuthHeader) < 2 {
				ape.RenderErr(w, problems.Unauthorized())
				return
			}
			claims, err := helpers.RetrieveClaimsFromJwtString(splitAuthHeader[1], secret)
			if err != nil {
				ape.RenderErr(w, problems.BadRequest(err)...)
				return
			}

			splitModulePermission := strings.Split(claims.ModulePermission, "/")

			permissionMap := make(map[string]string)
			for _, modulePermission := range splitModulePermission {
				split := strings.Split(modulePermission, ".")
				if len(split) < 2 {
					continue
				}

				permissionMap[split[0]] = split[1]
			}

			for _, permission := range permissions {
				if permissionMap[module] == permission {
					next.ServeHTTP(w, r)
					return
				}
			}

			ape.RenderErr(w, problems.Forbidden())
		})
	}
}
