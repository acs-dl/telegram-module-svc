package helpers

import (
	"fmt"

	"github.com/acs-dl/auth-svc/internal/data"
)

func CreatePermissionsString(permissions []data.ModulePermission) (string, error) {
	var resultPermission string

	for _, permission := range permissions {
		resultPermission += fmt.Sprintf("%s.%s/", permission.ModuleName, permission.PermissionName)
	}

	return resultPermission, nil
}
