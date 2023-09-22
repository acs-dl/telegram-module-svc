package data

type Chats interface {
	New() Chats

	Upsert(chat Chat) error
	Delete() error
	Get() (*Chat, error)
	Select() ([]Chat, error)

	FilterByTitles(titles ...string) Chats
	FilterByIds(ids ...int64) Chats
	FilterByAccessHash(accessHash *int64) Chats

	SearchBy(title string) Chats
}

type Chat struct {
	Title         string  `json:"title" db:"title" structs:"title"`
	Id            int64   `json:"id" db:"id" structs:"id"`
	AccessHash    *int64  `json:"access_hash" db:"access_hash" structs:"access_hash"`
	MembersAmount int64   `json:"members_amount" db:"members_amount" structs:"members_amount"`
	PhotoName     *string `json:"photo_name" db:"photo_name" structs:"photo_name"`
	PhotoLink     *string `json:"photo_link" db:"photo_link" structs:"photo_link"`
}
