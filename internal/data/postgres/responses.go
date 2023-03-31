package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const responsesTableName = "responses"

type ResponsesQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

var selectedResponsesTable = sq.Select("*").From(responsesTableName)

func NewResponsesQ(db *pgdb.DB) data.Responses {
	return &ResponsesQ{
		db:  db.Clone(),
		sql: selectedResponsesTable,
	}
}

func (q *ResponsesQ) New() data.Responses {
	return NewResponsesQ(q.db)
}

func (q *ResponsesQ) Insert(response data.Response) error {
	clauses := structs.Map(response)

	query := sq.Insert(responsesTableName).SetMap(clauses)

	return q.db.Exec(query)
}

func (q *ResponsesQ) Select() ([]data.Response, error) {
	var result []data.Response

	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *ResponsesQ) Get() (*data.Response, error) {
	var result data.Response

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *ResponsesQ) Delete(id string) error {
	query := sq.Delete(responsesTableName).Where(
		sq.Eq{"id": id})

	result, err := q.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.New("no responses with such id")
	}

	return nil
}

func (q *ResponsesQ) FilterById(id string) data.Responses {
	q.sql = q.sql.Where(sq.Eq{"id": id})

	return q
}

func (q *ResponsesQ) ResetFilters() data.Responses {
	q.sql = selectedResponsesTable

	return q
}
