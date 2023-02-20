package requests

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetUnverifiedUsersRequest struct {
	pgdb.OffsetPageParams

	Username *string `filter:"username"`
}

func NewGetUnverifiedUsersRequest(r *http.Request) (GetUnverifiedUsersRequest, error) {
	var request GetUnverifiedUsersRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
