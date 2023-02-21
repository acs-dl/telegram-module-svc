package handlers

import (
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/models"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/service/api/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func GetPermissions(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetPermissionsRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if request.UserId != nil && request.Link == nil {
		permissions, err := PermissionsQ(r).WithUsers().FilterByUserIds(*request.UserId).Page(request.OffsetPageParams).Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to get permissions")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		PermissionsQ(r).ResetFilters()

		totalCount, err := PermissionsQ(r).CountWithUsers().FilterByUserIds(*request.UserId).GetTotalCount()
		if err != nil {
			Log(r).WithError(err).Error("failed to get permissions total count")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		response := models.NewUserPermissionListResponse(permissions)
		response.Meta.TotalCount = totalCount
		response.Links = data.GetOffsetLinksForPGParams(r, request.OffsetPageParams)

		ape.Render(w, response)
		return
	}

	var userIds []int64
	if request.UserId != nil {
		userIds = append(userIds, *request.UserId)
	}

	var link = ""
	if request.Link != nil {
		link = *request.Link
	}

	permissions, err := PermissionsQ(r).WithUsers().FilterByUserIds(userIds...).SearchBy(link).Page(request.OffsetPageParams).Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get permissions")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	PermissionsQ(r).ResetFilters()

	totalCount, err := PermissionsQ(r).CountWithUsers().FilterByUserIds(userIds...).SearchBy(link).GetTotalCount()
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
