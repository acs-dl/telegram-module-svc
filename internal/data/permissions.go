package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	Owner  = "owner"
	Admin  = "admin"
	Member = "member"
	Left   = "left"
	Self   = "self"
	Banned = "banned"
)

type Permissions interface {
	New() Permissions

	Create(permission Permission) error
	Upsert(permission Permission) error
	UpdateAccessLevel(permission Permission) error
	Delete(telegramId int64, link string) error

	Select() ([]Permission, error)
	Get() (*Permission, error)

	FilterByTelegramIds(telegramIds ...int64) Permissions
	FilterByLinks(links ...string) Permissions
	FilterByTime(time time.Time) Permissions
	SearchBy(search string) Permissions

	WithUsers() Permissions
	FilterByUserIds(userIds ...int64) Permissions

	Count() Permissions
	CountWithUsers() Permissions
	GetTotalCount() (int64, error)

	Page(pageParams pgdb.OffsetPageParams) Permissions
}

type Permission struct {
	RequestId   string    `json:"request_id" db:"request_id" structs:"request_id"`
	TelegramId  int64     `json:"telegram_id" db:"telegram_id" structs:"telegram_id"`
	AccessLevel string    `json:"access_level" db:"access_level" structs:"access_level"`
	Link        string    `json:"link" db:"link" structs:"link"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" structs:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" structs:"-"`
	*User       `structs:",omitempty"`
}
