package models

import (
	"fmt"

	"github.com/acs-dl/telegram-module-svc/resources"
)

func newRole(name string, value string) resources.Role {
	return resources.Role{
		Key: resources.Key{
			ID:   fmt.Sprintf("%s-%s", name, value),
			Type: resources.ROLE,
		},
		Attributes: resources.RoleAttributes{
			Name:  name,
			Value: value,
		},
	}
}

func NewRoleResponse(name string, value string) resources.RoleResponse {
	return resources.RoleResponse{
		Data: newRole(name, value),
	}
}
