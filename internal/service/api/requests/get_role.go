package requests

import (
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetRoleRequest struct {
	AccessLevel *string `filter:"accessLevel"`
}

func NewGetRoleRequest(r *http.Request) (GetRoleRequest, error) {
	var request GetRoleRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
