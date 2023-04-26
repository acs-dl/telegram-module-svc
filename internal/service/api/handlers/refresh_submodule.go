package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/worker"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RefreshSubmodule(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRefreshSubmoduleRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	workerInstance := *worker.WorkerInstance(ParentContext(r.Context()))
	for _, link := range request.Data.Attributes.Links {
		go func(linkToHandle string) {
			err = workerInstance.CreatePermissions(linkToHandle)
			if err != nil {
				panic(err)
			}
		}(link)
	}

	w.WriteHeader(http.StatusAccepted)
	ape.Render(w, http.StatusAccepted)
}
