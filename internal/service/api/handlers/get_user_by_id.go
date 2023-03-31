package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := requests.NewGetUserByIdRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, err := UsersQ(r).FilterById(&userId).Get()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get user with id `%d`", userId)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if user == nil {
		Log(r).Errorf("no user with id `%d`", userId)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	permission, err := PermissionsQ(r).FilterByTelegramIds(user.TelegramId).Get()
	if err != nil {
		Log(r).Errorf("failed to get submodule for user with id `%d`", userId)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if permission != nil {
		user.Submodule = &permission.Link
	}

	ape.Render(w, models.NewUserResponse(*user))
}
