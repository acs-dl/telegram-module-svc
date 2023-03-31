package ape

import (
	"context"
	"net/http"
)

var CtxMiddleware = CtxMiddleWare

// DEPRECATED: use typo-free CtxMiddleware instead
func CtxMiddleWare(extenders ...func(context.Context) context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			for _, extender := range extenders {
				ctx = extender(ctx)
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
