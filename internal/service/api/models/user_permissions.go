package models

import (
	"strconv"

	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/acs-dl/telegram-module-svc/resources"
)

func NewUserPermissionModel(permission data.Permission, counter int) resources.UserPermission {
	if permission.Phone != nil {
		phoneWithoutCode := (*permission.Phone)[3:]
		permission.Phone = &phoneWithoutCode
	}

	id, accessHash := helpers.SubmoduleIdentifiersToString(permission.SubmoduleId, permission.SubmoduleAccessHash)

	result := resources.UserPermission{
		Key: resources.Key{
			ID:   strconv.Itoa(counter),
			Type: resources.USER_PERMISSION,
		},
		Attributes: resources.UserPermissionAttributes{
			Username:            permission.Username,
			Phone:               permission.Phone,
			ModuleId:            permission.TelegramId,
			UserId:              permission.Id,
			Link:                permission.Link,
			SubmoduleId:         id,
			SubmoduleAccessHash: accessHash,
			Path:                permission.Link,
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
