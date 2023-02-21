package handlers

import (
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

// TODO: think about roles
// when we add user ALWAYS member
// update can be up to admin or back yo member (only for owner <- it's a problem)
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

	if request.Username != nil {
		//GET user by phone of username
		user, err := UsersQ(r).FilterByUsernames(*request.Username).Get()
		if err != nil {
			Log(r).WithError(err).Infof("failed to get user with `%s` username", *request.Username)
			ape.RenderErr(w, problems.InternalError())
		}
		if user == nil {
			Log(r).WithError(err).Infof("no user was found with `%s` username", *request.Username)
			ape.RenderErr(w, problems.InternalError())
		}

		permission, err := PermissionsQ(r).FilterByTelegramIds(user.TelegramId).FilterByLinks(*request.Link).Get()
		if err != nil {
			Log(r).WithError(err).Infof("failed to get permission from `%s` to `%s`", *request.Link, *request.Username)
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if permission != nil {
			ape.Render(w, models.NewRolesResponse(true, permission.AccessLevel))
			return
		}
	}

	if request.Username != nil {
		chatUser, err := tg.NewTg(Params(r), Log(r)).GetChatUserFromApi(request.Username, nil, *request.Link)
		if err != nil {
			Log(r).WithError(err).Info("failed to check user from api")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		if chatUser != nil {
			ape.Render(w, models.NewRolesResponse(true, chatUser.AccessLevel))
			return
		}
	}

	ape.Render(w, models.NewRolesResponse(true, ""))
}
