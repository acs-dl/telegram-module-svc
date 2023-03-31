package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetUsersRequest struct {
	Username *string `filter:"username"`
	Phone    *string `filter:"phone"`
}

func NewGetUsersRequest(r *http.Request) (GetUsersRequest, error) {
	var request GetUsersRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
