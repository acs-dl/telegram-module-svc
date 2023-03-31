package data

type PermissionUsers interface {
	New() PermissionUsers

	Create(PermissionUser PermissionUser) (PermissionUser, error)
	Select() ([]PermissionUser, error)
	Get() (*PermissionUser, error)
	Delete(permission PermissionUser) error

	FilterByUserId(userId int64) PermissionUsers
	FilterByPermissionId(permissionId int64) PermissionUsers
}

type PermissionUser struct {
	PermissionId int64 `db:"permission_id" structs:"permission_id"`
	UserId       int64 `db:"user_id" structs:"user_id"`
}
