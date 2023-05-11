package data

type UserStatus string

const (
	SUPER_ADMIN UserStatus = "super_admin"
	ADMIN       UserStatus = "admin"
	USER        UserStatus = "user"
)

type Users interface {
	New() Users

	Get() (*User, error)

	FilterByEmails(emails ...string) Users
	FilterByIds(id ...int64) Users
}

type User struct {
	Id       int64      `db:"id" structs:"-"`
	Email    string     `db:"email" structs:"email"`
	Password string     `db:"password" structs:"password"`
	Status   UserStatus `db:"status" structs:"status"`
}
