package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/tokatu4561/go-ddd-demo/infrastructure/transaction"
)

var TxKey = struct{}{}

type tx struct {
    db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) transaction.Transaction {
    return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) error) (error) {
    tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
    if err != nil {
        return err
    }

    ctx = context.WithValue(ctx, &TxKey, tx)

    err = f(ctx)
    if err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Commit(); err != nil {
        // エラーならロールバック
        tx.Rollback()
        return err
    }

    return nil
}