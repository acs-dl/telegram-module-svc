package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type GetRolesRequest struct {
	Link                *string `filter:"link"`
	SubmoduleId         *string `filter:"submodule_id"`
	SubmoduleAccessHash *string `filter:"submodule_access_hash"`
	Username            *string `filter:"username"`
	Phone               *string `filter:"phone"`
}

func NewGetRolesRequest(r *http.Request) (GetRolesRequest, error) {
	var request GetRolesRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
