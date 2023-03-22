package urlval

import (
	"fmt"
	"strings"
)

// errBadRequest is an error type that can be rendered to client application
// using ape package (https://gitlab.com/distributed_lab/ape).
// Mostly useful when have a need to render multiple errors:
//
// 	errs := errBadRequest{}
//  for cond {
// 		if cond {
//			errs["k"] = errors.New("something wrong with field 'k'")
//		}
//	}
// 	return errs.Filter()
//
type errBadRequest map[string]error

// Error returns a concatenated errors in string format.
func (e errBadRequest) Error() string {
	var msg strings.Builder
	var sep = ", "

	for f, err := range e {
		msg.WriteString(sep)
		msg.WriteString(fmt.Sprintf("%s: %s", f, err.Error()))
	}

	return strings.TrimPrefix(msg.String(), sep)
}

// Filter returns errBadRequest if it contains any errors, nil -
// otherwise. Do return errs.Filter() when populating errors under for loop.
func (e errBadRequest) Filter() error {
	for k, err := range e {
		if err == nil {
			delete(e, k)
		}
	}

	if len(e) == 0 {
		return nil
	}

	return e
}

// BadRequest returns a message to be rendered to client in
// case it's request is invalid.
func (e errBadRequest) BadRequest() map[string]error {
	return e
}
