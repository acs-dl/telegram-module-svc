package pgdb

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Opts struct {
	URL                string
	MaxOpenConnections int
	MaxIdleConnections int
}

func Open(opts Opts) (*DB, error) {
	db, err := sqlx.Connect("postgres", opts.URL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	db.SetMaxIdleConns(opts.MaxIdleConnections)
	db.SetMaxOpenConns(opts.MaxOpenConnections)
	return &DB{
		db:      db,
		Queryer: newQueryer(db),
	}, nil
}

type Execer interface {
	Exec(query squirrel.Sqlizer) error
	ExecContext(ctx context.Context, query squirrel.Sqlizer) error
	ExecRaw(query string, args ...interface{}) error
	ExecRawContext(ctx context.Context, query string, args ...interface{}) error
	ExecWithResult(query squirrel.Sqlizer) (sql.Result, error)
	ExecWithResultContext(ctx context.Context, query squirrel.Sqlizer) (sql.Result, error)
}

type Selecter interface {
	Select(dest interface{}, query squirrel.Sqlizer) error
	SelectContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error
	SelectRaw(dest interface{}, query string, args ...interface{}) error
	SelectRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Getter interface {
	Get(dest interface{}, query squirrel.Sqlizer) error
	GetContext(ctx context.Context, dest interface{}, query squirrel.Sqlizer) error
	GetRaw(dest interface{}, query string, args ...interface{}) error
	GetRawContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type TransactionFunc func() error

type Transactor interface {
	Transaction(transactionFunc TransactionFunc) (err error)
}

// Connection is yet another thin wrapper for sql.DB allowing to use squirrel queries directly
type Connection interface {
	Transactor
	Queryer
}

// Queryer overloads sqlx's interface name with different meaning, which is not cool.
type Queryer interface {
	Execer
	Selecter
	Getter
}
