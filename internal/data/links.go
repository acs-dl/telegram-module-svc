package data

type Links interface {
	New() Links

	Insert(link Link) error
	Delete(link string) error

	Get() (*Link, error)
	Select() ([]Link, error)
}

type Link struct {
	Id   int64  `db:"id" structs:"-"`
	Link string `db:"link" structs:"link"`
}
