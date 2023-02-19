package requests

import (
	"gitlab.com/distributed_lab/urlval"
	"net/http"
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
