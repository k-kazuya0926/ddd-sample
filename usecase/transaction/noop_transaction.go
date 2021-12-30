package transaction

import "context"

// 何もしない実装
type NoopTransaction struct {
}

func (nt *NoopTransaction) DoInTx(ctx context.Context, f func(context.Context) error) error {
	return f(ctx)
}
