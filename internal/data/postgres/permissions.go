package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/acs-dl/telegram-module-svc/internal/data"
	"github.com/acs-dl/telegram-module-svc/internal/helpers"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	permissionsTableName                 = "permissions"
	permissionsRequestIdColumn           = permissionsTableName + ".request_id"
	permissionsTelegramIdColumn          = permissionsTableName + ".telegram_id"
	permissionsLinkColumn                = permissionsTableName + ".link"
	permissionsSubmoduleIdColumn         = permissionsTableName + ".submodule_id"
	permissionsSubmoduleAccessHashColumn = permissionsTableName + ".submodule_access_hash"
	permissionsAccessLevelColumn         = permissionsTableName + ".access_level"
	permissionsCreatedAtColumn           = permissionsTableName + ".created_at"
	permissionsUpdatedAtColumn           = permissionsTableName + ".updated_at"
)

type PermissionsQ struct {
	db            *pgdb.DB
	selectBuilder sq.SelectBuilder
	deleteBuilder sq.DeleteBuilder
	updateBuilder sq.UpdateBuilder
}

var permissionsColumns = []string{
	permissionsRequestIdColumn,
	permissionsTelegramIdColumn,
	permissionsLinkColumn,
	permissionsAccessLevelColumn,
	permissionsCreatedAtColumn,
	permissionsUpdatedAtColumn,
	permissionsSubmoduleIdColumn,
	permissionsSubmoduleAccessHashColumn,
}

func NewPermissionsQ(db *pgdb.DB) data.Permissions {
	return &PermissionsQ{
		db:            db.Clone(),
		selectBuilder: sq.Select(permissionsColumns...).From(permissionsTableName),
		deleteBuilder: sq.Delete(permissionsTableName),
		updateBuilder: sq.Update(permissionsTableName),
	}
}

func (q PermissionsQ) New() data.Permissions {
	return NewPermissionsQ(q.db)
}

func (q PermissionsQ) UpdateAccessLevel(permission data.Permission) error {
	query := q.updateBuilder.Set("access_level", permission.AccessLevel)

	return q.db.Exec(query)
}

func (q PermissionsQ) Select() ([]data.Permission, error) {
	var result []data.Permission

	err := q.db.Select(&result, q.selectBuilder)

	return result, err
}

func (q PermissionsQ) Upsert(permission data.Permission) error {
	updateStmt, args := sq.Update(" ").
		Set("updated_at", time.Now()).
		Set("access_level", permission.AccessLevel).MustSql()

	query := sq.Insert(permissionsTableName).SetMap(structs.Map(permission)).
		Suffix("ON CONFLICT (telegram_id, submodule_id, submodule_access_hash) DO "+updateStmt, args...)

	return q.db.Exec(query)
}

func (q PermissionsQ) Get() (*data.Permission, error) {
	var result data.Permission

	err := q.db.Get(&result, q.selectBuilder)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q PermissionsQ) Delete() error {
	var deleted []data.Permission

	err := q.db.Select(&deleted, q.deleteBuilder.Suffix("RETURNING *"))
	if err != nil {
		return err
	}

	if len(deleted) == 0 {
		return errors.Errorf("no such data to delete")
	}

	return nil
}

func (q PermissionsQ) FilterByTelegramIds(telegramIds ...int64) data.Permissions {
	equalTelegramIds := sq.Eq{permissionsTelegramIdColumn: telegramIds}
	q.selectBuilder = q.selectBuilder.Where(equalTelegramIds)
	q.deleteBuilder = q.deleteBuilder.Where(equalTelegramIds)
	q.updateBuilder = q.updateBuilder.Where(equalTelegramIds)

	return q
}

func (q PermissionsQ) FilterByLinks(links ...string) data.Permissions {
	equalLinks := sq.Eq{permissionsLinkColumn: links}
	q.selectBuilder = q.selectBuilder.Where(equalLinks)
	q.deleteBuilder = q.deleteBuilder.Where(equalLinks)
	q.updateBuilder = q.updateBuilder.Where(equalLinks)

	return q
}

func (q PermissionsQ) SearchBy(search string) data.Permissions {
	search = strings.Replace(search, " ", "%", -1)
	search = fmt.Sprint("%", search, "%")
	ilikeSearch := sq.ILike{permissionsLinkColumn: search}

	q.selectBuilder = q.selectBuilder.Where(ilikeSearch)
	q.deleteBuilder = q.deleteBuilder.Where(ilikeSearch)
	q.updateBuilder = q.updateBuilder.Where(ilikeSearch)

	return q
}

func (q PermissionsQ) Count() data.Permissions {
	q.selectBuilder = sq.Select("COUNT (*)").From(permissionsTableName)

	return q
}

func (q PermissionsQ) GetTotalCount() (int64, error) {
	var count int64
	err := q.db.Get(&count, q.selectBuilder)

	return count, err
}

func (q PermissionsQ) Page(pageParams pgdb.OffsetPageParams) data.Permissions {
	q.selectBuilder = pageParams.ApplyTo(q.selectBuilder, "link")

	return q
}

func (q PermissionsQ) WithUsers() data.Permissions {
	q.selectBuilder = sq.Select().Columns(helpers.RemoveDuplicateColumn(append(permissionsColumns, usersColumns...))...).
		From(permissionsTableName).
		LeftJoin(usersTableName + " ON " + usersTelegramIdColumn + " = " + permissionsTelegramIdColumn).
		Where(sq.NotEq{permissionsRequestIdColumn: nil}).
		GroupBy(helpers.RemoveDuplicateColumn(append(permissionsColumns, usersColumns...))...)

	return q
}

func (q PermissionsQ) CountWithUsers() data.Permissions {
	q.selectBuilder = sq.Select("COUNT(*)").From(permissionsTableName).
		LeftJoin(usersTableName + " ON " + usersTelegramIdColumn + " = " + permissionsTelegramIdColumn).
		Where(sq.NotEq{permissionsRequestIdColumn: nil})

	return q
}

func (q PermissionsQ) FilterByUserIds(userIds ...int64) data.Permissions {
	equalUserIds := sq.Eq{usersIdColumn: userIds}

	if len(userIds) == 0 {
		equalUserIds = sq.Eq{usersIdColumn: nil}
	}

	q.selectBuilder = q.selectBuilder.Where(equalUserIds)
	q.deleteBuilder = q.deleteBuilder.Where(equalUserIds)
	q.updateBuilder = q.updateBuilder.Where(equalUserIds)

	return q
}

func (q PermissionsQ) FilterByGreaterTime(time time.Time) data.Permissions {
	greaterTime := sq.Gt{permissionsUpdatedAtColumn: time}

	q.selectBuilder = q.selectBuilder.Where(greaterTime)
	q.deleteBuilder = q.deleteBuilder.Where(greaterTime)
	q.updateBuilder = q.updateBuilder.Where(greaterTime)

	return q
}

func (q PermissionsQ) FilterByLowerTime(time time.Time) data.Permissions {
	lowerTime := sq.Lt{permissionsUpdatedAtColumn: time}

	q.selectBuilder = q.selectBuilder.Where(lowerTime)
	q.deleteBuilder = q.deleteBuilder.Where(lowerTime)
	q.updateBuilder = q.updateBuilder.Where(lowerTime)

	return q
}
