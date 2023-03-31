package data

type Modules interface {
	New() Modules

	Create(module Module) (*Module, error)
	Select() ([]Module, error)
	GetByName(name string) (*Module, error)
	Delete(moduleName string) error
}

type Module struct {
	Id   int64  `db:"id" structs:"-"`
	Name string `db:"name" structs:"name"`
}

type ModulePermission struct {
	Id             int64  `db:"id" structs:"-"`
	ModuleId       int64  `db:"module_id" structs:"module_id"`
	ModuleName     string `db:"module_name" structs:"name"`
	PermissionName string `db:"permission_name" structs:"name"`
}
