package postgres

import (
	"database/sql"
	"fmt"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const usersTableName = "users"

type GitlabUsersQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

var selectedUsersTable = sq.Select("*").From(usersTableName)

func NewUsersQ(db *pgdb.DB) data.Users {
	return &GitlabUsersQ{
		db:  db.Clone(),
		sql: selectedUsersTable,
	}
}

func (q *GitlabUsersQ) New() data.Users {
	return NewUsersQ(q.db)
}

func (q *GitlabUsersQ) Upsert(user data.User) error {
	clauses := structs.Map(user)

	stmt := "ON CONFLICT (gitlab_id) DO UPDATE SET created_at = CURRENT_TIMESTAMP"
	if user.Id != nil {
		stmt = fmt.Sprintf("ON CONFLICT (gitlab_id) DO UPDATE SET created_at = CURRENT_TIMESTAMP, id = %d", *user.Id)
	}
	query := sq.Insert(usersTableName).SetMap(clauses).Suffix(stmt)

	return q.db.Exec(query)
}

func (q *GitlabUsersQ) GetById(id int64) (*data.User, error) {
	query := q.sql.Where(sq.Eq{"id": id})

	var result data.User
	err := q.db.Get(&result, query)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *GitlabUsersQ) GetByUsername(username string) (*data.User, error) {
	query := q.sql.Where(sq.Eq{"username": username})

	var result data.User
	err := q.db.Get(&result, query)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *GitlabUsersQ) GetByGitlabId(gitlabId int64) (*data.User, error) {
	query := q.sql.Where(sq.Eq{"gitlab_id": gitlabId})

	var result data.User
	err := q.db.Get(&result, query)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *GitlabUsersQ) Delete(gitlabId int64) error {
	query := sq.Delete(usersTableName).Where(
		sq.Eq{"gitlab_id": gitlabId})

	result, err := q.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.Errorf("no users with id `%d`", gitlabId)
	}

	return nil
}

func (q *GitlabUsersQ) FilterByIds(ids ...*int64) data.Users {
	stmt := sq.Eq{usersTableName + ".id": nil}
	if ids == nil {
		stmt = sq.Eq{usersTableName + ".id": ids}
	}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *GitlabUsersQ) Get() (*data.User, error) {
	var result data.User

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *GitlabUsersQ) Select() ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *GitlabUsersQ) Page(pageParams pgdb.OffsetPageParams) data.Users {
	q.sql = pageParams.ApplyTo(q.sql, "username")

	return q
}

func (q *GitlabUsersQ) SearchBy(search string) data.Users {
	search = strings.Replace(search, " ", "%", -1)
	search = fmt.Sprint("%", search, "%")

	q.sql = q.sql.Where(sq.ILike{"username": search})
	return q
}

func (q *GitlabUsersQ) Count() data.Users {
	q.sql = sq.Select("COUNT (*)").From(usersTableName)

	return q
}

func (q *GitlabUsersQ) GetTotalCount() (int64, error) {
	var count int64
	err := q.db.Get(&count, q.sql)

	return count, err
}
