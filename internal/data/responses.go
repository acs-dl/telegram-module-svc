package data

type Responses interface {
	New() Responses

	Get() (*Response, error)
	Select() ([]Response, error)

	Insert(response Response) error
	Delete(id string) error

	FilterById(id string) Responses
	ResetFilters() Responses
}

type Response struct {
	ID     string `json:"id" db:"id" structs:"id"`
	Status string `json:"status" db:"status" structs:"status"`
	Error  string `json:"error" db:"error" structs:"error"`
	//Payload   json.RawMessage `json:"payload" db:"payload"`
	//CreatedAt string          `json:"created_at" db:"created_at" structs:"-"`
}
