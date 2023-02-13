package ape

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"gitlab.com/distributed_lab/logan/v3"
)

// LoggerSetter - sets specified entry into ctx
type LoggerSetter func(ctx context.Context, entry *logan.Entry) context.Context

// RequestIDProvider - returns ID of the request from ctx if one present
type RequestIDProvider func(context.Context) string

// LoganMiddleware by default will just log the begging and end of the request
// Also it might take use of number of optionally injected arguments like:
// - time.Duration will set threshold for logging a warning of slow request
// - LoggerSetter - will be called with *logan.Entry which will contain auxiliary fields to identify the request
// - RequestIDProvider - will set result into fields of *logan.Entry if LoggerSetter is provided
func LoganMiddleware(entry *logan.Entry, args ...interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			entry := entry.WithField("path", r.URL.Path)
			requestID := tryGetRequestID(r.Context(), args)
			if requestID != "" {
				entry = entry.WithField("request_id", requestID)
			}

			r = tryCallLoggerSetter(r, entry, args)
			durationThreshold := tryGetDuration(args)
			defer func() {
				dur := time.Since(start)
				entry = entry.WithFields(logan.F{
					"duration": dur,
					"status":   ww.Status(),
				})
				entry.Info("request finished")

				if durationThreshold != nil && dur > *durationThreshold {
					entry.WithField("http_request", r).Warn("slow request")
				}
			}()

			entry.WithField("path", r.URL.Path).Info("request started")
			next.ServeHTTP(ww, r)
		})
	}
}

func tryGetDuration(args []interface{}) *time.Duration {
	for i := range args {
		switch v := args[i].(type) {
		case time.Duration:
			return &v
		}
	}

	return nil
}

func tryCallLoggerSetter(request *http.Request, entry *logan.Entry, args []interface{}) *http.Request {
	for i := range args {
		switch v := args[i].(type) {
		case LoggerSetter:
			newCtx := v(request.Context(), entry)
			return request.WithContext(newCtx)
		}
	}

	return request
}

func tryGetRequestID(ctx context.Context, args []interface{}) string {
	for i := range args {
		switch v := args[i].(type) {
		case RequestIDProvider:
			return v(ctx)
		}
	}

	return ""
}
