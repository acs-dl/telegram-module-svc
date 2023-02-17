package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type Users interface {
	New() Users

	Upsert(user User) error
	Delete(telegramId int64) error

	Select() ([]User, error)
	Get() (*User, error)

	FilterById(id *int64) Users
	FilterByTelegramIds(telegramIds ...int64) Users
	FilterByUsernames(usernames ...string) Users
	FilterByPhones(phones ...string) Users
	SearchBy(search string) Users

	ResetFilters() Users

	Count() Users
	GetTotalCount() (int64, error)

	Page(pageParams pgdb.OffsetPageParams) Users
}

type User struct {
	Id         *int64    `json:"-" db:"id" structs:"id,omitempty"`
	Username   string    `json:"username" db:"username" structs:"username"`
	Phone      string    `json:"phone" db:"phone" structs:"phone"`
	FirstName  string    `json:"first_name" db:"first_name" structs:"first_name"`
	LastName   string    `json:"last_name" db:"last_name" structs:"last_name"`
	TelegramId int64     `json:"telegram_id" db:"telegram_id" structs:"telegram_id"`
	AccessHash int64     `json:"access_hash" db:"access_hash" structs:"access_hash"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" structs:"created_at"`
	// fields to create permission
	AccessLevel string `json:"-" db:"-" structs:"-"`
}
