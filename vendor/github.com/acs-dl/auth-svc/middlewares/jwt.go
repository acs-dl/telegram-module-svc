package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/acs-dl/auth-svc/internal/service/api/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Jwt(secret, module string, permissions ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Println("no auth-svc header provided")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			splitAuthHeader := strings.Split(authHeader, " ")
			if len(splitAuthHeader) != 2 {
				log.Println("header must consist from two parts")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}
			claims, err := helpers.RetrieveClaimsFromJwtString(splitAuthHeader[1], secret)
			if err != nil {
				log.Println(err, ":failed to retrieve claims from jwt string")
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

			log.Println("no needed permission was granted")
			ape.RenderErr(w, problems.Forbidden())
		})
	}
}
