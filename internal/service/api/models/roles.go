package models

import (
	"github.com/acs-dl/telegram-module-svc/resources"
)

var roles = []resources.AccessLevel{
	{Name: "Member", Value: "member"},
}

func NewRolesModel(found bool, roles []resources.AccessLevel, chats []resources.Chat) resources.Roles {
	result := resources.Roles{
		Key: resources.Key{
			ID:   "roles",
			Type: resources.ROLES,
		},
		Attributes: resources.RolesAttributes{
			Req:   found,
			Roles: roles,
			Chats: chats,
		},
	}

	return result
}

func NewRolesResponse(found bool, chats []resources.Chat) resources.RolesResponse {
	return resources.RolesResponse{
		Data: NewRolesModel(found, roles, chats),
	}
}
