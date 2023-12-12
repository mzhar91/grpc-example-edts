package repository

import (
	"context"
	"database/sql"
	"net/http"
)

type Use struct {
	Db    *sql.DB
	Trans Transaction
}

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Trans func(Transaction) (error, int)

func WithTransaction(db *sql.DB, trans Trans) (error, int) {
	errCode := 0
	tx, err := db.Begin()
	if err != nil {
		return err, http.StatusInternalServerError
	}

	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			if err != nil {
				// logrus.Fatal(err)
			}

			return
		} else if err != nil {
			err = tx.Rollback()
			if err != nil {
				// logrus.Fatal(err)
			}

			return
		} else {
			err = tx.Commit()
			if err != nil {
				// logrus.Fatal(err)
			}

			return
		}
	}()

	err, errCode = trans(tx)

	return err, errCode
}
