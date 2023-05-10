package data

import "encoding/json"

const (
	SUCCESS = "success"
	FAILURE = "failure"
	INVITED = "invited"
)

type Responses interface {
	New() Responses

	Insert(response Response) error
	Get() (*Response, error)
	Select() ([]Response, error)
	Delete() error

	FilterByIds(ids ...string) Responses
}

type Response struct {
	ID          string          `json:"id" db:"id" structs:"id"`
	Status      string          `json:"status" db:"status" structs:"status"`
	Error       *string         `json:"error,omitempty" db:"error,omitempty" structs:"error,omitempty"`
	Description *string         `json:"description,omitempty" db:"description,omitempty" structs:"description,omitempty"`
	Payload     json.RawMessage `json:"payload" db:"payload" structs:"payload"`
	CreatedAt   string          `json:"created_at" db:"created_at" structs:"-"`
}
