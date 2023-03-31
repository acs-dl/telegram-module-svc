package middleware

import (
	"fmt"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

// TODO: get current program host to compare with URL host
func CheckHost(programHost string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestHost := r.Host
			fmt.Println(requestHost)
			if requestHost == programHost {
				next.ServeHTTP(w, r)
				return
			}

			ape.RenderErr(w, problems.Forbidden())
		})
	}
}
