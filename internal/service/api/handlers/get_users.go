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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetUsersRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	username := ""
	if request.Username != nil {
		username = *request.Username
	}

	users, err := UsersQ(r).SearchBy(username).Select()
	if err != nil {
		Log(r).WithError(err).Errorf("failed to select users from db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(users) != 0 {
		ape.Render(w, models.NewUserInfoListResponse(users, 0))
		return
	}

	users, err = searchUsersFromApi(PQueues(ParentContext(r.Context())).SuperPQueue, tg.NewTg(Params(r), Log(r)), request.Username, request.Phone, 10)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to search users from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewUserInfoListResponse(users, 0))
}

func searchUsersFromApi(pq *pqueue.PriorityQueue, tgClient tg.TelegramClient, username, phone *string, amount int) ([]data.User, error) {
	item, err := helpers.AddFunctionInPQueue(pq, tgClient.SearchByFromApi, []any{any(username), any(phone), any(amount)}, pqueue.HighPriority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add function in pqueue")
	}

	err = item.Response.Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to check user from api")
	}

	users, ok := item.Response.Value.([]data.User)
	if !ok {
		return nil, errors.Wrap(err, "wrong user type in response")
	}

	return users, nil
}
