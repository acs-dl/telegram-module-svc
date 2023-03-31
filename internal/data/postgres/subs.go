package postgres

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/acs/telegram-module/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
)

const subsTableName = "subs"

type SubsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

var (
	subsColumns       = []string{subsTableName + ".id", subsTableName + ".link as subs_link", subsTableName + ".path", subsTableName + ".lpath", subsTableName + ".type as subs_type", subsTableName + ".parent_id"}
	selectedSubsTable = sq.Select(subsColumns...).From(subsTableName)
)

func NewSubsQ(db *pgdb.DB) data.Subs {
	return &SubsQ{
		db:  db.Clone(),
		sql: selectedSubsTable,
	}
}

func (q *SubsQ) New() data.Subs {
	return NewSubsQ(q.db)
}

func (q *SubsQ) Insert(sub data.Sub) error {
	clauses := structs.Map(sub)

	query := sq.Insert(subsTableName).SetMap(clauses)

	return q.db.Exec(query)
}

func (q *SubsQ) Upsert(sub data.Sub) error {
	query := sq.Insert(subsTableName).SetMap(structs.Map(sub)).
		Suffix(fmt.Sprintf("ON CONFLICT DO NOTHING"))

	return q.db.Exec(query)
}

func (q *SubsQ) Select() ([]data.Sub, error) {
	var result []data.Sub

	err := q.db.Select(&result, q.sql)

	return result, err
}

func (q *SubsQ) DistinctOn(column string) data.Subs {
	q.sql = q.sql.Options(fmt.Sprintf("DISTINCT ON (%s)", column))

	return q
}

func (q *SubsQ) Get() (*data.Sub, error) {
	var result data.Sub

	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *SubsQ) Delete(subId int64, typeTo, link string) error {
	query := sq.Delete(subsTableName).Where(
		sq.Eq{"id": subId, "type": typeTo, "link": link})

	result, err := q.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	affectedRows, _ := result.RowsAffected()
	if affectedRows == 0 {
		return errors.New("no sub with such data")
	}

	return nil
}

func (q *SubsQ) FilterByParentIds(parentIds ...int64) data.Subs {
	stmt := sq.Eq{subsTableName + ".parent_id": parentIds}
	if len(parentIds) == 0 {
		stmt = sq.Eq{subsTableName + ".parent_id": nil}
	}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *SubsQ) FilterByLinks(links ...string) data.Subs {
	stmt := sq.Eq{subsTableName + ".link": links}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *SubsQ) FilterByIds(ids ...int64) data.Subs {
	stmt := sq.Eq{subsTableName + ".id": ids}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *SubsQ) FilterByUserIds(userIds ...int64) data.Subs {
	stmt := sq.Eq{permissionsTableName + ".user_id": userIds}

	if len(userIds) == 0 {
		stmt = sq.Eq{permissionsTableName + ".user_id": nil}
	}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *SubsQ) FilterByGitlabIds(gitlabIds ...int64) data.Subs {
	stmt := sq.Eq{permissionsTableName + ".gitlab_id": gitlabIds}

	q.sql = q.sql.Where(stmt)

	return q
}

func (q *SubsQ) ResetFilters() data.Subs {
	q.sql = selectedSubsTable

	return q
}

func (q *SubsQ) FilterByLevel(lpath ...string) data.Subs {
	query := sq.Expr(fmt.Sprintf("subs.lpath ?? ARRAY[%s]::lquery[]", strings.Join(lpath, ",")))

	q.sql = q.sql.Where(query).
		OrderBy("subs.lpath").PlaceholderFormat(sq.Dollar)
	return q
}

func (q *SubsQ) FilterByLowerLevel(parentLpath string) data.Subs {
	q.sql = q.sql.Where(fmt.Sprintf("%s.lpath <@ '%s'", subsTableName, parentLpath))

	return q
}

func (q *SubsQ) FilterExceptSelf(parentLpath string) data.Subs {
	q.sql = q.sql.Where(fmt.Sprintf("%s.lpath <> '%s'", subsTableName, parentLpath))

	return q
}

func (q *SubsQ) FilterByHigherLevel(parentLpath string) data.Subs {
	q.sql = q.sql.Where(fmt.Sprintf("%s.lpath @> '%s'", subsTableName, parentLpath))

	return q
}

func (q *SubsQ) OrderBy(columns ...string) data.Subs {
	q.sql = q.sql.OrderBy(columns...)

	return q
}

func (q *SubsQ) WithPermissions() data.Subs {
	q.sql = sq.Select().Columns(subsColumns...).Column(fmt.Sprintf("nlevel(%s.lpath) as level", subsTableName)).
		Columns(permissionsTableName+".request_id", permissionsTableName+".user_id", permissionsTableName+".name", permissionsTableName+".username", permissionsTableName+".gitlab_id", permissionsTableName+".access_level").
		From(subsTableName).
		LeftJoin(fmt.Sprint(permissionsTableName, " ON ", permissionsTableName, ".link = ", subsTableName, ".link")).
		Where(sq.NotEq{permissionsTableName + ".request_id": nil})

	return q
}

func (q *SubsQ) CountWithPermissions() data.Subs {
	q.sql = sq.Select("COUNT(*)").From(subsTableName).
		LeftJoin(fmt.Sprint(permissionsTableName, " ON ", permissionsTableName, ".link = ", subsTableName, ".link")).
		Where(sq.NotEq{permissionsTableName + ".request_id": nil})

	return q
}

func (q *SubsQ) SearchBy(search string) data.Subs {
	search = strings.Replace(search, " ", "%", -1)
	search = fmt.Sprint("%", search, "%")

	q.sql = q.sql.Where(sq.ILike{subsTableName + ".link": search})

	return q
}

func (q *SubsQ) FilterByParentLevel(parentLevel bool) data.Subs {
	q.sql = q.sql.Where(sq.Eq{permissionsTableName + ".parent_level": parentLevel})

	return q
}

func (q *SubsQ) Count() data.Subs {
	q.sql = sq.Select("COUNT (*)").From(subsTableName)

	return q
}

func (q *SubsQ) GetTotalCount() (int64, error) {
	var count int64
	stmt, args, _ := q.sql.ToSql()
	fmt.Println(stmt, args)
	err := q.db.Get(&count, q.sql)

	return count, err
}

func (q *SubsQ) Page(pageParams pgdb.OffsetPageParams) data.Subs {
	q.sql = pageParams.ApplyTo(q.sql, subsTableName+".link")

	return q
}
