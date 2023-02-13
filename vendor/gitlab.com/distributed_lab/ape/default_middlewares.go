package ape

import (
	"errors"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"
)

// DefaultMiddlewares provide sane defaults that should just work
func DefaultMiddlewares(r chi.Router, args ...interface{}) {
	log := tryFindLogan(args)
	if log == nil {
		panic(errors.New("*logan.Entry is required"))
	}

	// push new args to front to ensure that override any existing
	args = append([]interface{}{
		LoggerSetter(SetContextLog),
		LoggerGetter(getContextLog),
		RequestIDProvider(GetRequestID)}, args...)
	r.Use(
		CtxMiddleware(setRequestID),
		// log middleware goes first, so subsequent entries will have request context
		LoganMiddleware(log, args...),
		RecoverMiddleware(args...),
		ContentType("application/json"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := problems.NotFound()
		err.Detail = "Unknown path"
		err.Meta = &map[string]interface{}{
			"url": r.URL.String(),
		}
		RenderErr(w, err)
	})
}

func tryFindLogan(args []interface{}) *logan.Entry {
	for _, arg := range args {
		switch v := arg.(type) {
		case *logan.Entry:
			return v
		}
	}

	return nil
}
