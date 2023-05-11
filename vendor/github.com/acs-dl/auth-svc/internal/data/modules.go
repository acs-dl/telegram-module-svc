package data

type Modules interface {
	New() Modules

	Insert(module Module) error
	Select() ([]Module, error)
	Get() (*Module, error)
	Delete() error

	FilterByNames(names ...string) Modules
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
