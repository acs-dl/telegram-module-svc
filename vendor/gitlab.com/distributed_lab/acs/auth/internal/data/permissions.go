package data

type Permissions interface {
	New() Permissions

	Create(module Permission) (*Permission, error)
	Select() ([]ModulePermission, error)
	Get() (*ModulePermission, error)
	Delete(permission Permission) error

	WithModules() Permissions

	FilterByModuleName(name string) Permissions
	FilterByPermissionId(permissionId int64) Permissions

	ResetFilters() Permissions
}

type Permission struct {
	Id       int64  `db:"id" structs:"-"`
	ModuleId int64  `db:"module_id" structs:"module_id"`
	Name     string `db:"name" structs:"name"`
	*Module  `db:"-" structs:",omitempty"`
}
