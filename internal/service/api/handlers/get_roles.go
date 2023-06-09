package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/requests"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
	"github.com/acs-dl/telegram-module-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetRolesRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Link == nil {
		ape.Render(w, models.NewRolesResponse(false, make([]resources.Chat, 0)))
		return
	}

	username := ""
	if request.Username != nil {
		username = *request.Username
	}

	phone := ""
	if request.Phone != nil {
		phone = *request.Phone
		if phone[0:1] == "+" {
			phone = phone[1:]
		}
	}

	user, err := UsersQ(r).FilterByUsername(username).FilterByPhone(phone).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get user with `%s` username and `%s` phone", username, phone)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if user != nil {
		permission, err := PermissionsQ(r).FilterByTelegramIds(user.TelegramId).FilterByLinks(*request.Link).Get()
		if err != nil {
			Log(r).WithError(err).Errorf("failed to get permission from `%s` to `%s`/`%s`", *request.Link, username, phone)
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if permission != nil {
			Log(r).Warnf("user `%s`/`%s` is already in `%s`", username, phone, *request.Link)
			ape.RenderErr(w, problems.Conflict())
			return
		}
	}

	pqs := pqueue.PQueuesInstance(ParentContext(r.Context()))
	tgClient := tg_client.TelegramClientInstance(ParentContext(r.Context()))

	user, err = helpers.GetUser(
		pqs.UserPQueue,
		any(tgClient.GetUserFromApi),
		[]any{
			any(tgClient.GetUsualClient()),
			any(request.Username),
			any(&phone),
		},
		pqueue.HighPriority,
	)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get user from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	chats, err := helpers.GetChats(pqs.SuperUserPQueue, tgClient.GetChatFromApi, []any{any(*request.Link)}, pqueue.HighPriority)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chat from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	chat := chats[0]

	chatUser, err := helpers.GetUser(pqs.SuperUserPQueue, tgClient.GetChatUserFromApi, []any{any(*user), any(chat)}, pqueue.HighPriority)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chat user from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if chatUser != nil {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	dbChats, err := ChatsQ(r).FilterByTitles(*request.Link).Select()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chats with `%s` title", *request.Link)
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// when we add user ALWAYS member
	ape.Render(w, models.NewRolesResponse(true, models.NewChatListModel(dbChats)))
}
