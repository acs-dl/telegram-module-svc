package data

type Users interface {
	New() Users

	Get() (*User, error)

	FilterByEmail(email string) Users
	FilterById(id int64) Users
}

type User struct {
	Id       int64  `db:"id" structs:"-"`
	Email    string `db:"email" structs:"email"`
	Password string `db:"password" structs:"password"`
}
