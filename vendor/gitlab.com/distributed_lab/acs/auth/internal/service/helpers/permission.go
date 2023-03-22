package helpers

import (
	"fmt"
	"gitlab.com/distributed_lab/acs/auth/internal/data"
)

func CreatePermissionsString(permissions []data.PermissionUser, PermissionsQ data.Permissions) (string, error) {
	var resultPermission string

	for _, permission := range permissions {
		module, err := PermissionsQ.WithModules().FilterByPermissionId(permission.PermissionId).Get()
		if err != nil {
			return "", err
		}
		resultPermission += fmt.Sprintf("%s.%s/", module.ModuleName, module.PermissionName)
		PermissionsQ.ResetFilters()
	}
	return resultPermission, nil
}
