package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/helpers"
	"gitlab.com/distributed_lab/logan/v3/errors"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	usersTableName        = "users"
	usersIdColumn         = usersTableName + ".id"
	usersUsernameColumn   = usersTableName + ".username"
	usersPhoneColumn      = usersTableName + ".phone"
	usersTelegramIdColumn = usersTableName + ".telegram_id"
	usersAccessHashColumn = usersTableName + ".access_hash"
	usersFirstNameColumn  = usersTableName + ".first_name"
	usersLastNameColumn   = usersTableName + ".last_name"
	usersCreatedAtColumn  = usersTableName + ".created_at"
	usersUpdatedAtColumn  = usersTableName + ".updated_at"
)

type UsersQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
	updateBuilder sq.UpdateBuilder
}

var (
	usersColumns = []string{
		usersIdColumn,
		usersUsernameColumn,
		usersPhoneColumn,
		usersTelegramIdColumn,
		usersAccessHashColumn,
		usersFirstNameColumn,
		usersLastNameColumn,
		usersCreatedAtColumn,
	}
	selectedUsersTable = sq.Select("*").From(usersTableName)
)

func NewUsersQ(db *pgdb.DB) data.Users {
	return &UsersQ{
		db:            db.Clone(),
		selectBuilder: selectedUsersTable,
		deleteBuilder: sq.Delete(usersTableName),
		updateBuilder: sq.Update(usersTableName),
	}
}

func (q UsersQ) New() data.Users {
	return NewUsersQ(q.db)
}

func (q UsersQ) Upsert(user data.User) error {
	if user.Phone != nil && *user.Phone == "" {
		user.Phone = nil
	}
	if user.Username != nil && *user.Username == "" {
		user.Username = nil
	}

	clauses := structs.Map(user)

	updateQuery := sq.Update(" ").
		Set("updated_at", time.Now())

	if user.Id != nil {
		updateQuery = updateQuery.Set("id", *user.Id)
	}

	updateStmt, args := updateQuery.MustSql()

	query := sq.Insert(usersTableName).SetMap(clauses).Suffix("ON CONFLICT (telegram_id) DO "+updateStmt, args...)

	return q.db.Exec(query)
}

func (q UsersQ) Delete() error {
	var deleted []data.User

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return errors.Errorf("no such data to delete")
	}

	return nil
}

func (q UsersQ) Get() (*data.User, error) {
	var result data.User

	err := q.db.Get(&result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q UsersQ) Select() ([]data.User, error) {
	var result []data.User

	err := q.db.Select(&result, q.selectBuilder)

	return result, err
}

func (q UsersQ) FilterById(id *int64) data.Users {
	equalId := sq.Eq{usersIdColumn: id}

	q.selectBuilder = q.selectBuilder.Where(equalId)
	q.deleteBuilder = q.deleteBuilder.Where(equalId)
	q.updateBuilder = q.updateBuilder.Where(equalId)

	return q
}

func (q UsersQ) FilterByTelegramIds(telegramIds ...int64) data.Users {
	equalTelegramIds := sq.Eq{usersTelegramIdColumn: telegramIds}

	q.selectBuilder = q.selectBuilder.Where(equalTelegramIds)
	q.deleteBuilder = q.deleteBuilder.Where(equalTelegramIds)
	q.updateBuilder = q.updateBuilder.Where(equalTelegramIds)

	return q
}

func (q UsersQ) FilterByUsername(username string) data.Users {
	if err := helpers.ValidateNonEmptyString(username); err != nil {
		return q

	}
	equalUsername := sq.Eq{usersUsernameColumn: username}

	q.selectBuilder = q.selectBuilder.Where(equalUsername)
	q.deleteBuilder = q.deleteBuilder.Where(equalUsername)
	q.updateBuilder = q.updateBuilder.Where(equalUsername)

	return q
}

func (q UsersQ) FilterByPhone(phone string) data.Users {
	if err := helpers.ValidateNonEmptyString(phone); err != nil {
		return q

	}

	equalPhone := sq.Eq{usersPhoneColumn: phone}

	q.selectBuilder = q.selectBuilder.Where(equalPhone)
	q.deleteBuilder = q.deleteBuilder.Where(equalPhone)
	q.updateBuilder = q.updateBuilder.Where(equalPhone)

	return q
}

func (q UsersQ) Page(pageParams pgdb.OffsetPageParams) data.Users {
	q.selectBuilder = pageParams.ApplyTo(q.selectBuilder, "username")

	return q
}

func (q UsersQ) SearchBy(search string) data.Users {
	search = strings.Replace(search, " ", "%", -1)
	search = fmt.Sprint("%", search, "%")

	q.selectBuilder = q.selectBuilder.Where(sq.ILike{usersUsernameColumn: search})

	return q
}

func (q UsersQ) Count() data.Users {
	q.selectBuilder = sq.Select("COUNT (*)").From(usersTableName)

	return q
}

func (q UsersQ) GetTotalCount() (int64, error) {
	var count int64
	err := q.db.Get(&count, q.selectBuilder)

	return count, err
}

func (q UsersQ) FilterByLowerTime(time time.Time) data.Users {
	lowerTime := sq.Lt{usersUpdatedAtColumn: time}

	q.selectBuilder = q.selectBuilder.Where(lowerTime)
	q.deleteBuilder = q.deleteBuilder.Where(lowerTime)
	q.updateBuilder = q.updateBuilder.Where(lowerTime)

	return q
}
