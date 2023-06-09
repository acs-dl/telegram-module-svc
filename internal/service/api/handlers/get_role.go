package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetRole(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetRoleRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.AccessLevel == nil {
		Log(r).Errorf("no access level was provided")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	name := data.Roles[*request.AccessLevel]
	if name == "" {
		Log(r).Errorf("no such access level `%s`", *request.AccessLevel)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	ape.Render(w, models.NewRoleResponse(data.Roles[*request.AccessLevel], *request.AccessLevel))
}
