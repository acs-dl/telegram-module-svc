package pgdb

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type rawQueryer interface {
	sqlx.ExecerContext
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type queryer struct {
	raw rawQueryer
}

func newQueryer(raw rawQueryer) *queryer {
	return &queryer{
		raw: raw,
	}
}

func (q *queryer) Get(dest interface{}, query squirrel.Sqlizer) error {
	return q.GetContext(context.Background(), dest, query)
}

func (q *queryer) GetContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error {
	sql, args, err := build(query)
	if err != nil {
		return err
	}
	return q.GetRawContext(ctx, dest, sql, args...)
}

func (q *queryer) GetRaw(dest interface{}, query string, args ...interface{}) error {
	return q.GetRawContext(context.Background(), dest, query, args...)
}

func (q *queryer) GetRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	query = rebind(query)
	err := q.raw.GetContext(ctx, dest, query, args...)

	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to get raw")
}

func (q *queryer) Exec(query squirrel.Sqlizer) error {
	return q.ExecContext(context.Background(), query)
}

func (q *queryer) ExecContext(ctx context.Context, query squirrel.Sqlizer) error {
	sql, args, err := build(query)
	if err != nil {
		return errors.Wrap(err, "failed to build query")
	}
	return q.ExecRawContext(ctx, sql, args...)
}

func (q *queryer) ExecRaw(query string, args ...interface{}) error {
	return q.ExecRawContext(context.Background(), query, args...)
}

func (q *queryer) ExecRawContext(ctx context.Context, query string, args ...interface{}) error {
	query = rebind(query)
	_, err := q.raw.ExecContext(ctx, query, args...)
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to exec query")
}

func (q *queryer) ExecWithResult(query squirrel.Sqlizer) (sql.Result, error) {
	return q.ExecWithResultContext(context.Background(), query)
}

func (q *queryer) ExecWithResultContext(ctx context.Context, query squirrel.Sqlizer) (sql.Result, error) {
	sql, args, err := build(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build query")
	}
	return q.ExecRawWithResultContext(ctx, sql, args...)
}

func (q *queryer) ExecRawWithResult(query string, args ...interface{}) (sql.Result, error) {
	return q.ExecRawWithResultContext(context.Background(), query, args...)
}

func (q *queryer) ExecRawWithResultContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	query = rebind(query)
	result, err := q.raw.ExecContext(ctx, query, args...)
	if err == nil {
		return result, nil
	}

	if err == sql.ErrNoRows {
		return nil, err
	}

	return nil, errors.Wrap(err, "failed to exec query")
}

// Select runs `query`, setting the results found on `dest`.
func (q *queryer) Select(dest interface{}, query squirrel.Sqlizer) error {
	return q.SelectContext(context.Background(), dest, query)
}

// SelectContext runs `query`, setting the results found on `dest`.
func (q *queryer) SelectContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error {
	sql, args, err := build(query)
	if err != nil {
		return err
	}
	return q.SelectRawContext(ctx, dest, sql, args...)
}

func (q *queryer) SelectRaw(dest interface{}, query string, args ...interface{}) error {
	return q.SelectRawContext(context.Background(), dest, query, args...)
}

// SelectRawContext runs `query` with `args`, setting the results found on `dest`.
func (q *queryer) SelectRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	//r.clearSliceIfPossible(dest) // TODO wat?
	query = rebind(query)
	err := q.raw.SelectContext(ctx, dest, query, args...)

	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to select")
}

func rebind(stmt string) string {
	return sqlx.Rebind(sqlx.BindType("postgres"), stmt)
}

func build(b squirrel.Sqlizer) (sql string, args []interface{}, err error) {
	sql, args, err = b.ToSql()

	if err != nil {
		err = errors.Wrap(err, "failed to parse query")
	}
	return
}
