package models

import (
	"strconv"

	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

func NewUserPermissionModel(permission data.Permission, counter int) resources.UserPermission {
	result := resources.UserPermission{
		Key: resources.Key{
			ID:   strconv.Itoa(counter),
			Type: resources.USER_PERMISSION,
		},
		Attributes: resources.UserPermissionAttributes{
			Username: permission.Username,
			Phone:    permission.Phone,
			ModuleId: permission.User.TelegramId,
			UserId:   permission.Id,
			Link:     permission.Link,
			Path:     permission.Link,
			AccessLevel: resources.AccessLevel{
				Name:  data.Roles[permission.AccessLevel],
				Value: permission.AccessLevel,
			},
		},
	}

	return result
}

func NewUserPermissionList(permissions []data.Permission) []resources.UserPermission {
	result := make([]resources.UserPermission, len(permissions))
	for i, permission := range permissions {
		result[i] = NewUserPermissionModel(permission, i)
	}
	return result
}

func NewUserPermissionListResponse(permissions []data.Permission) UserPermissionListResponse {
	return UserPermissionListResponse{
		Data: NewUserPermissionList(permissions),
	}
}

type UserPermissionListResponse struct {
	Meta  Meta                       `json:"meta"`
	Data  []resources.UserPermission `json:"data"`
	Links *resources.Links           `json:"links"`
}
