package transaction

import (
	"context"
	"ddd-sample/usecase/transaction"
)

// 何もしない実装
type noopTransaction struct {
}

func NewNoopTransaction() transaction.Transaction {
	return &noopTransaction{}
}

func (nt *noopTransaction) DoInTx(ctx context.Context, f func(context.Context) error) error {
	return f(ctx)
}
