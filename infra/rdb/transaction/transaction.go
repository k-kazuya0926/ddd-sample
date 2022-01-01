package transaction

import (
	"context"
	"database/sql"
	"ddd-sample/usecase/transaction"

	"github.com/jmoiron/sqlx"
)

var txKey = struct{}{}

type tx struct {
	db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) transaction.Transaction {
	return &tx{db: db}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	// ここでctxへトランザクションオブジェクトを放り込む。
	ctx = context.WithValue(ctx, &txKey, tx)

	// トランザクションの対象処理へコンテキストを引き継ぎ
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

// context.Contextからトランザクションを取得する関数
func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}
