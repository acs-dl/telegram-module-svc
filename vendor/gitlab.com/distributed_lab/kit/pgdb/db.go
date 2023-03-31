package pgdb

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DB struct {
	Queryer
	db *sqlx.DB
}

func (db *DB) RawDB() *sql.DB {
	return db.db.DB
}

func (db *DB) Clone() *DB {
	return &DB{
		Queryer: newQueryer(db.db),
		db:      db.db,
	}

}

// Transaction is generic helper method for specific Q's to implement Transaction capabilities
func (db *DB) Transaction(fn TransactionFunc) (err error) {
	return db.TransactionWithOptions(nil, fn)
}

func (db *DB) TransactionWithOptions(opts *sql.TxOptions, fn TransactionFunc) (err error) {
	// TODO panic on nested tx

	tx, err := db.db.BeginTxx(context.TODO(), opts)
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}

	db.Queryer = newQueryer(tx)

	// swallowing rollback err, should not affect data consistency
	defer tx.Rollback()
	defer func() {
		db.Queryer = newQueryer(db.db)
	}()

	if err = fn(); err != nil {
		return errors.Wrap(err, "failed to execute statements")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return
}

