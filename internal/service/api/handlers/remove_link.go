package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func RemoveLink(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewRemoveLinkRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse remove link request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = LinksQ(r).Delete(request.Data.Attributes.Link)
	if err != nil {
		Log(r).WithError(err).Error("failed to delete link")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	Log(r).Infof("successfully removed link `%s`", request.Data.Attributes.Link)
	w.WriteHeader(http.StatusAccepted)
	ape.Render(w, http.StatusAccepted)
}
