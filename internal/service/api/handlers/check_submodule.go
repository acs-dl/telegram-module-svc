package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CheckSubmodule(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCheckSubmoduleRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Link == nil {
		Log(r).Errorf("no link was provided")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	chats, err := ChatsQ(r).SearchBy(*request.Link).Select()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chats with `%s` title", *request.Link)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(chats) == 0 {
		Log(r).Warnf("no chats were found")
		ape.Render(w, models.NewLinkResponse("", false, chats))
		return
	}

	ape.Render(w, models.NewLinkResponse(*request.Link, true, chats))
}
