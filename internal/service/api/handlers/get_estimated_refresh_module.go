package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/pqueue"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"
	"gitlab.com/distributed_lab/ape"
)

func GetEstimatedRefreshModule(w http.ResponseWriter, r *http.Request) {
	parentContext := ParentContext(r.Context())
	workerInstance := *worker.WorkerInstance(parentContext)

	pqueueRequestsAmount := int64(pqueue.PQueuesInstance(parentContext).SuperPQueue.Len() + pqueue.PQueuesInstance(parentContext).UsualPQueue.Len())
	requestsTimeLimit := Config(parentContext).RateLimit().TimeLimit
	requestsAmountLimit := Config(parentContext).RateLimit().RequestsAmount

	timeToHandleOneRequest := requestsTimeLimit / time.Duration(requestsAmountLimit)
	estimatedTime := time.Duration(pqueueRequestsAmount)*timeToHandleOneRequest + workerInstance.GetEstimatedTime()

	ape.Render(w, models.NewEstimatedTimeResponse(estimatedTime.String()))
}
