package handlers

import (
	"net/http"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/models"
	"github.com/acs-dl/telegram-module-svc/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetPermissions(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPermissionsRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	var userIds []int64
	if request.UserId != nil {
		userIds = append(userIds, *request.UserId)
	}

	permissionsQ := PermissionsQ(r).WithUsers().FilterByUserIds(userIds...)
	countPermissionsQ := PermissionsQ(r).CountWithUsers().FilterByUserIds(userIds...)

	if request.Link != nil {
		permissionsQ = permissionsQ.SearchBy(*request.Link)
		countPermissionsQ = countPermissionsQ.SearchBy(*request.Link)
	}

	permissions, err := permissionsQ.Page(request.OffsetPageParams).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get permissions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	totalCount, err := countPermissionsQ.GetTotalCount()
	if err != nil {
		Log(r).WithError(err).Error("failed to get permissions total count")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := models.NewUserPermissionListResponse(permissions)
	response.Meta.TotalCount = totalCount
	response.Links = data.GetOffsetLinksForPGParams(r, request.OffsetPageParams)

	ape.Render(w, response)
}
