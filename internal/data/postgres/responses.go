package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	responsesTableName = "responses"
	responsesIdColumn  = responsesTableName + ".id"
)

type ResponsesQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
}

var selectedResponsesTable = sq.Select("*").From(responsesTableName)

func NewResponsesQ(db *pgdb.DB) data.Responses {
	return &ResponsesQ{
		db:            db.Clone(),
		selectBuilder: selectedResponsesTable,
		deleteBuilder: sq.Delete(responsesTableName),
	}
}

func (q ResponsesQ) New() data.Responses {
	return NewResponsesQ(q.db)
}

func (q ResponsesQ) Insert(response data.Response) error {
	clauses := structs.Map(response)

	query := sq.Insert(responsesTableName).SetMap(clauses)

	return q.db.Exec(query)
}

func (q ResponsesQ) Select() ([]data.Response, error) {
	var result []data.Response

	err := q.db.Select(&result, q.selectBuilder)

	return result, err
}

func (q ResponsesQ) Get() (*data.Response, error) {
	var result data.Response

	err := q.db.Get(&result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q ResponsesQ) Delete() error {
	var deleted []data.Response

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return errors.Errorf("no such data to delete")
	}

	return nil
}

func (q ResponsesQ) FilterByIds(ids ...string) data.Responses {
	equalIds := sq.Eq{responsesIdColumn: ids}

	q.selectBuilder = q.selectBuilder.Where(equalIds)
	q.deleteBuilder = q.deleteBuilder.Where(equalIds)

	return q
}
