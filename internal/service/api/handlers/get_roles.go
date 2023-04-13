package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetRolesRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.Link == nil {
		ape.Render(w, models.NewRolesResponse(false, ""))
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
		Log(r).WithError(err).Infof("failed to get user with `%s` username and `%s` phone", username, phone)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if user != nil {
		permission, err := PermissionsQ(r).FilterByTelegramIds(user.TelegramId).FilterByLinks(*request.Link).Get()
		if err != nil {
			Log(r).WithError(err).Infof("failed to get permission from `%s` to `%s`/`%s`", *request.Link, username, phone)
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if permission != nil {
			ape.Render(w, models.NewRolesResponse(true, permission.AccessLevel))
			return
		}
	}

	chatUser, err := getUserFromApi(PQueues(ParentContext(r.Context())).SuperPQueue, tg.NewTg(Params(r), Log(r)), request.Username, request.Phone, *request.Link)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get user from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if chatUser != nil {
		ape.Render(w, models.NewRolesResponse(true, chatUser.AccessLevel))
		return
	}

	// when we add user ALWAYS member
	ape.Render(w, models.NewRolesResponse(true, data.Admin))
}

func getUserFromApi(pq *pqueue.PriorityQueue, tgClient tg.TelegramClient, username, phone *string, link string) (*data.User, error) {
	item, err := helpers.AddFunctionInPQueue(pq, tgClient.GetChatUserFromApi, []any{any(username), any(phone), any(link)}, pqueue.HighPriority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to check user from api")
	}

	chatUser, ok := item.Response.Value.(*data.User)
	if !ok {
		return nil, errors.Wrap(err, "wrong user type in response")
	}

	return chatUser, nil
}
