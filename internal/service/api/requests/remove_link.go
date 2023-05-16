package requests

import (
	"encoding/json"
	"net/http"

	"github.com/acs-dl/telegram-module-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type RemoveLinkRequest struct {
	Data resources.Link `json:"data"`
}

func NewRemoveLinkRequest(r *http.Request) (RemoveLinkRequest, error) {
	var request RemoveLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, " failed to unmarshal")
	}

	return request, request.validate()
}

func (r *RemoveLinkRequest) validate() error {
	return validation.Errors{
		"link": validation.Validate(&r.Data.Attributes.Link, validation.Required),
	}.Filter()
}
