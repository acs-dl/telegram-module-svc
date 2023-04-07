package data

type Links interface {
	New() Links

	Insert(link Link) error
	Delete() error
	Get() (*Link, error)
	Select() ([]Link, error)

	FilterByLinks(links ...string) Links
}

type Link struct {
	Id   int64  `db:"id" structs:"-"`
	Link string `db:"link" structs:"link"`
}
