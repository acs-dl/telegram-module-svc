package ape

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getsentry/raven-go"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

// Logger abstract logger interface providing all possible bells and whistles,
// mainly for decoupling `distributed_lab/logan` package
type Logger interface {
	Log(level uint32, fields map[string]interface{}, err error, withStack bool, args ...interface{})
}

// LoggerGetter get *logan.Entry from request context
type LoggerGetter func(ctx context.Context) *logan.Entry

// RecoverMiddleware by default will just catch handler panic.
// Also it might take use of number of optionally injected arguments like:
// - LoggerGetter request context log getter
// - Logger implementation to log stacktrace with Error level
// - *jsonapi.ErrorObject to render error body
// - *raven.Client to report exception to Sentry
func RecoverMiddleware(args ...interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					responseRendered := false
					rerr := errors.FromPanic(rvr)
					for _, arg := range args {
						switch v := arg.(type) {
						case Logger:
							// 2 - is logan/logrus error level as it most probable implementation of `Logger`
							v.Log(2, nil, rerr, true, "handler panicked")
						case LoggerGetter:
							v(r.Context()).WithStack(rerr).WithError(rerr).Error("handler panicked")
						case *jsonapi.ErrorObject:
							RenderErr(w, v)
							responseRendered = true
						case *raven.Client:
							packet := raven.NewPacket(
								fmt.Sprint(rvr),
								raven.NewException(
									rerr,
									raven.GetOrNewStacktrace(rerr, 0, 0, nil)),
								raven.NewHttp(r))
							v.Capture(packet, nil)
						}
					}
					if !responseRendered {
						RenderErr(w, problems.InternalError())
					}
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
