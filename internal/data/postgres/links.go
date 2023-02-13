package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const linksTableName = "links"

type LinksQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
}

func NewLinksQ(db *pgdb.DB) data.Links {
	return &LinksQ{
		db:            db,
		selectBuilder: sq.Select(linksTableName + ".*").From(linksTableName),
	}
}

func (r *LinksQ) New() data.Links {
	return NewLinksQ(r.db)
}

func (r *LinksQ) FilterByLinks(links ...string) data.Links {
	stmt := sq.Eq{linksTableName + ".link": links}
	r.selectBuilder = r.selectBuilder.Where(stmt)
	return r
}

func (r *LinksQ) Get() (*data.Link, error) {
	var result data.Link
	err := r.db.Get(&result, r.selectBuilder)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *LinksQ) Select() ([]data.Link, error) {
	var result []data.Link

	err := r.db.Select(&result, r.selectBuilder)

	return result, errors.Wrap(err, "failed to select links")
}

func (r *LinksQ) Insert(link data.Link) error {
	insertStmt := sq.Insert(linksTableName).SetMap(structs.Map(link)).Suffix("ON CONFLICT (link) DO NOTHING")
	err := r.db.Exec(insertStmt)
	return errors.Wrap(err, "failed to insert link")
}

func (r *LinksQ) Delete(link string) error {
	query := sq.Delete(linksTableName).Where(
		sq.Eq{"link": link})

	result, err := r.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.New("no such link")
	}

	return nil
}
