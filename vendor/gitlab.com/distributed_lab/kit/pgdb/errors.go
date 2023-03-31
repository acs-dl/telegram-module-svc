package pgdb

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func IsConstraintErr(err error, name string) bool {
	if err == nil {
		return false
	}
	cause := errors.Cause(err)
	pqerr, ok := cause.(*pq.Error)
	if !ok {
		return false
	}
	return pqerr.Constraint == name
}
