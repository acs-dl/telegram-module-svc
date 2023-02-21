package postgres

import (
	"database/sql"
	"fmt"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
	"time"

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
	usersTableName + ".id",
	usersTableName + ".username",
	usersTableName + ".phone",
	usersTableName + ".telegram_id",
	usersTableName + ".access_hash",
	usersTableName + ".first_name",
	usersTableName + ".last_name",
	usersTableName + ".created_at",
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
	if user.Phone != nil && *user.Phone == "" {
		user.Phone = nil
	}
	if user.Username != nil && *user.Username == "" {
		user.Username = nil
	}

	clauses := structs.Map(user)

	updateQuery := sq.Update(" ").
		Set("created_at", time.Now())

	if user.Id != nil {
		updateQuery = updateQuery.Set("id", *user.Id)
	}

	updateStmt, args := updateQuery.MustSql()

	query := sq.Insert(usersTableName).SetMap(clauses).Suffix("ON CONFLICT (telegram_id) DO "+updateStmt, args...)

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

func (q *UsersQ) FilterById(id *int64) data.Users {
	stmt := sq.Eq{usersTableName + ".id": id}

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
