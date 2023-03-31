package models

import (
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/resources"
)

var roles = []resources.AccessLevel{
	{Name: "Member", Value: "member"},
	{Name: "Admin", Value: "admin"},
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

func NewRolesResponse(found bool, current string) resources.RolesResponse {
	if !found {
		return resources.RolesResponse{
			Data: NewRolesModel(found, []resources.AccessLevel{}),
		}
	}

	newRoles := newRolesArray(current)
	if len(newRoles) == 0 {
		return resources.RolesResponse{
			Data: NewRolesModel(!found, []resources.AccessLevel{}),
		}
	}

	return resources.RolesResponse{
		Data: NewRolesModel(found, newRoles),
	}
}

func newRolesArray(current string) []resources.AccessLevel {
	result := make([]resources.AccessLevel, 0)

	if current == data.Owner {
		return result
	}

	for _, role := range roles {
		if role.Value != current {
			result = append(result, role)
		}
	}

	return result
}
