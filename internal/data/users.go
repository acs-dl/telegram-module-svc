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

	FilterByTime(time time.Time) Users
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
	Username   *string   `json:"username" db:"username" structs:"username,omitempty"`
	Phone      *string   `json:"phone" db:"phone" structs:"phone,omitempty"`
	FirstName  string    `json:"first_name" db:"first_name" structs:"first_name"`
	LastName   string    `json:"last_name" db:"last_name" structs:"last_name"`
	TelegramId int64     `json:"telegram_id" db:"telegram_id" structs:"telegram_id"`
	AccessHash int64     `json:"access_hash" db:"access_hash" structs:"access_hash"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" structs:"created_at"`
	Module     string    `json:"module" db:"-" structs:"-"`
	// fields to create permission
	AccessLevel string `json:"-" db:"-" structs:"-"`
}

type UnverifiedUser struct {
	CreatedAt time.Time `json:"created_at"`
	Module    string    `json:"module"`
	ModuleId  int64     `json:"module_id"`
	Email     *string   `json:"email,omitempty"`
	Name      *string   `json:"name,omitempty"`
	Phone     *string   `json:"phone,omitempty"`
	Username  *string   `json:"username,omitempty"`
}
