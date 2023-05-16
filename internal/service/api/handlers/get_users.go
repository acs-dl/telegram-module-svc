package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/requests"
	"github.com/acs-dl/telegram-module-svc/internal/tg_client"
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

	pq := pqueue.PQueuesInstance(ParentContext(r.Context())).SuperUserPQueue
	tgClient := tg_client.TelegramClientInstance(ParentContext(r.Context()))

	users, err = helpers.GetUsers(pq, tgClient.SearchByFromApi, []any{any(request.Username), any(request.Phone), any(10)}, pqueue.HighPriority)
	if err != nil {
		Log(r).WithError(err).Errorf("failed to get chat user from api")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewUserInfoListResponse(users, 0))
}
