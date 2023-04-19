package handlers

import (
	"math"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetEstimatedRefreshSubmodule(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRefreshSubmoduleRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	permissionsAmount, err := PermissionsQ(r).Count().FilterByLinks(request.Data.Attributes.Links...).GetTotalCount()
	if err != nil {
		Log(r).WithError(err).Error("failed to get permissions amount")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	parentContext := ParentContext(r.Context())

	pqueueRequestsAmount := int64(pqueue.PQueuesInstance(parentContext).SuperPQueue.Len() + pqueue.PQueuesInstance(parentContext).UsualPQueue.Len())

	requestsTimeLimit := Config(parentContext).RateLimit().TimeLimit
	requestsAmountLimit := Config(parentContext).RateLimit().RequestsAmount

	timeToHandleOneRequest := requestsTimeLimit / time.Duration(requestsAmountLimit)
	totalRequestsAmount := math.Round(float64(permissionsAmount+pqueueRequestsAmount) * 1.4)

	estimatedTime := time.Duration(totalRequestsAmount) * timeToHandleOneRequest

	ape.Render(w, models.NewEstimatedTimeResponse(estimatedTime.String()))
}
