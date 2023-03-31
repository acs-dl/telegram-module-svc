package requests

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewGetUserByIdRequest(r *http.Request) (int64, error) {
	id := chi.URLParam(r, "id")

	if id == "" {
		return 0, errors.New("`id` param is not specified")
	}

	return strconv.ParseInt(id, 10, 64)
}
