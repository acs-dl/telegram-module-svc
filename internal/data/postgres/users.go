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

type UsersQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

var selectedUsersTable = sq.Select("*").From(usersTableName)

var usersColumns = []string{
	permissionsTableName + ".id",
	permissionsTableName + ".username",
	permissionsTableName + ".phone",
	permissionsTableName + ".telegram_id",
	permissionsTableName + ".access_hash",
	permissionsTableName + ".first_name",
	permissionsTableName + ".last_name",
	permissionsTableName + ".created_at",
}

func NewUsersQ(db *pgdb.DB) data.Users {
	return &UsersQ{
		db:  db.Clone(),
		sql: selectedUsersTable,
	}
}

func (q *UsersQ) New() data.Users {
	return NewUsersQ(q.db)
}

func (q *UsersQ) Upsert(user data.User) error {
	clauses := structs.Map(user)

	stmt := "ON CONFLICT (telegram_id) DO UPDATE SET created_at = CURRENT_TIMESTAMP"
	if user.Id != nil {
		stmt = fmt.Sprintf("ON CONFLICT (gitlab_id) DO UPDATE SET created_at = CURRENT_TIMESTAMP, id = %d", *user.Id)
	}
	query := sq.Insert(usersTableName).SetMap(clauses).Suffix(stmt)

	return q.db.Exec(query)
}

func (q *UsersQ) GetById(id int64) (*data.User, error) {
	query := q.sql.Where(sq.Eq{"id": id})

	var result data.User
	err := q.db.Get(&result, query)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *UsersQ) Delete(telegramId int64) error {
	query := sq.Delete(usersTableName).Where(
		sq.Eq{"telegram_id": telegramId})

	result, err := q.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.Errorf("no users with id `%d`", telegramId)
	}

	return nil
}

func (q *UsersQ) Get() (*data.User, error) {
	var result data.User

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *UsersQ) Select() ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *UsersQ) FilterByIds(ids ...*int64) data.Users {
	stmt := sq.Eq{usersTableName + ".id": nil}
	if ids != nil {
		stmt = sq.Eq{usersTableName + ".id": ids}
	}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *UsersQ) FilterByTelegramIds(telegramIds ...int64) data.Users {
	q.sql = q.sql.Where(sq.Eq{usersTableName + ".telegram_id": telegramIds})

	return q
}

func (q *UsersQ) FilterByUsernames(usernames ...string) data.Users {
	q.sql = q.sql.Where(sq.Eq{usersTableName + ".username": usernames})

	return q
}

func (q *UsersQ) FilterByPhones(phones ...string) data.Users {
	q.sql = q.sql.Where(sq.Eq{usersTableName + ".phone": phones})

	return q
}

func (q *UsersQ) Page(pageParams pgdb.OffsetPageParams) data.Users {
	q.sql = pageParams.ApplyTo(q.sql, "username")

	return q
}

func (q *UsersQ) SearchBy(search string) data.Users {
	search = strings.Replace(search, " ", "%", -1)
	search = fmt.Sprint("%", search, "%")

	q.sql = q.sql.Where(sq.ILike{"username": search})
	return q
}

func (q *UsersQ) Count() data.Users {
	q.sql = sq.Select("COUNT (*)").From(usersTableName)

	return q
}

func (q *UsersQ) GetTotalCount() (int64, error) {
	var count int64
	err := q.db.Get(&count, q.sql)

	return count, err
}

func (q *UsersQ) ResetFilters() data.Users {
	q.sql = selectedResponsesTable

	return q
}
