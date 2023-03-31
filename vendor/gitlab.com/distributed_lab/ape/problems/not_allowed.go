package problems

import (
	"fmt"
	"net/http"

	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

func isBadRequest(err error) bool {
	e, ok := err.(interface {
		BadRequest() bool
	})
	return ok && e.BadRequest()
}

func isNotAllowed(err error) bool {
	e, ok := err.(interface {
		NotAllowed() bool
	})
	return ok && e.NotAllowed()
}

func isForbidden(err error) bool {
	e, ok := err.(interface {
		Forbidden() bool
	})
	return ok && e.Forbidden()
}

// NotAllowed will try to guess details of error and populate problem accordingly.
func NotAllowed(errs ...error) *jsonapi.ErrorObject {
	// errs is optional for backward compatibility
	if len(errs) == 0 {
		return &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusUnauthorized),
			Status: fmt.Sprintf("%d", http.StatusUnauthorized),
		}
	}

	if len(errs) != 1 {
		panic(errors.New("unexpected number of errors passed"))
	}

	cause := errors.Cause(errs[0])
	switch {
	case isBadRequest(cause):
		return &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "Request signature was invalid in some way",
			Meta: &map[string]interface{}{
				"reason": cause.Error(),
			},
		}
	case isNotAllowed(cause):
		return &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusUnauthorized),
			Status: fmt.Sprintf("%d", http.StatusUnauthorized),
		}
	case isForbidden(cause): {
		return &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusForbidden),
			Status: fmt.Sprintf("%d", http.StatusForbidden),
		}
	}
	default:
		return &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusInternalServerError),
			Status: fmt.Sprintf("%d", http.StatusInternalServerError),
		}
	}
}
