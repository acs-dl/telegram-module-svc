package handlers

import (
	"net/http"
	"time"

	"github.com/acs-dl/telegram-module-svc/internal/pqueue"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/worker"
	"gitlab.com/distributed_lab/ape"
)

func GetEstimatedRefreshModule(w http.ResponseWriter, r *http.Request) {
	parentContext := ParentContext(r.Context())
	workerInstance := *worker.WorkerInstance(parentContext)

	pqueueRequestsAmount := int64(pqueue.PQueuesInstance(parentContext).SuperUserPQueue.Len() + pqueue.PQueuesInstance(parentContext).UserPQueue.Len())
	requestsTimeLimit := Config(parentContext).RateLimit().TimeLimit
	requestsAmountLimit := Config(parentContext).RateLimit().RequestsAmount

	timeToHandleOneRequest := requestsTimeLimit / time.Duration(requestsAmountLimit)
	estimatedTime := time.Duration(pqueueRequestsAmount)*timeToHandleOneRequest + workerInstance.GetEstimatedTime()

	ape.Render(w, models.NewEstimatedTimeResponse(estimatedTime.String()))
}
