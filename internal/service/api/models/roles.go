package models

import (
	"github.com/acs-dl/telegram-module-svc/resources"
)

var Roles = []resources.AccessLevel{
	{Name: "Member", Value: "member"},
}

func NewRolesModel(found bool, roles []resources.AccessLevel) resources.Roles {
	result := resources.Roles{
		Key: resources.Key{
			ID:   "roles",
			Type: resources.ROLES,
		},
		Attributes: resources.RolesAttributes{
			Req:  found,
			List: roles,
		},
	}

	return result
}

func NewRolesResponse(found bool) resources.RolesResponse {
	roles := make([]resources.AccessLevel, 0)
	if found {
		roles = Roles
	}

	return resources.RolesResponse{
		Data: NewRolesModel(found, roles),
	}
}
