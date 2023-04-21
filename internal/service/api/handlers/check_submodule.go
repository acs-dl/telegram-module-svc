package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg_client"
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

	link, err := LinksQ(r).FilterByLinks(*request.Link).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get link `%s`", *request.Link)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if link != nil {
		ape.Render(w, models.NewLinkResponse(link.Link, true))
		return
	}

	telegramClient := tg_client.TelegramClientInstance(ParentContext(r.Context()))

	chat, err := helpers.GetChat(pqueue.PQueuesInstance(ParentContext(r.Context())).SuperPQueue, any(telegramClient.GetChatFromApi), []any{any(*request.Link)}, pqueue.HighPriority)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chat from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if chat != nil {
		ape.Render(w, models.NewLinkResponse(*request.Link, true))
		return
	}

	Log(r).Warnf("no group/project was found")
	ape.Render(w, models.NewLinkResponse("", false))
}
