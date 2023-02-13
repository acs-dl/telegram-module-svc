package ape

import (
	"fmt"
	"net/http"
	"time"
)

// CacheMiddleware returns middleware with directive for a browser to cache the response for some TTL (time to live).
func CacheMiddleware(ttl time.Duration) func(http.Handler) http.Handler {
	header := fmt.Sprintf("public, max-age=%d", ttl/time.Second)
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			// Adds directive before handling, as the handler may want to set its own cache control.
			rw.Header().Add("Cache-Control", header)
			handler.ServeHTTP(rw, r)
		})
	}
}
