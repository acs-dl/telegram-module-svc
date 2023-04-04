package handlers

import (
	"container/heap"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/tg"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
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

	newUuid := uuid.New()
	queueItem := &pqueue.QueueItem{
		Uuid:     newUuid,
		Func:     tg.NewTg(Params(r), Log(r)).SearchByFromApi,
		Args:     []any{any(request.Username), any(request.Phone), any(10)},
		Priority: 10,
	}
	heap.Push(PQueue(r.Context()), queueItem)
	item := PQueue(r.Context()).WaitUntilInvoked(newUuid)
	PQueue(r.Context()).RemoveByUUID(newUuid)
	err = item.Response.Error
	if err != nil {
		Log(r).WithError(err).Infof("failed to get users from api by `%s`", username)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	users, ok := item.Response.Value.([]data.User)
	if !ok {
		Log(r).WithError(err).Infof("wrong users type in response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewUserInfoListResponse(users, 0))
}
