package data

type Permissions interface {
	New() Permissions

	Insert(permission Permission) error
	Select() ([]ModulePermission, error)
	Get() (*ModulePermission, error)
	Delete() error

	IncludeModules() Permissions

	FilterByStatus(status UserStatus) Permissions
}

type Permission struct {
	Id       int64      `db:"id" structs:"-"`
	ModuleId int64      `db:"module_id" structs:"module_id"`
	Name     string     `db:"name" structs:"name"`
	Status   UserStatus `db:"status" structs:"status"`
	*Module  `db:"-" structs:",omitempty"`
}
