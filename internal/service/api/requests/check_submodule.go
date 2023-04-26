package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/urlval"
)

type CheckSubmoduleRequest struct {
	Link *string `filter:"link"`
}

func NewCheckSubmoduleRequest(r *http.Request) (CheckSubmoduleRequest, error) {
	var request CheckSubmoduleRequest

	err := urlval.Decode(r.URL.Query(), &request)

	return request, err
}
