//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package transaction

import "context"

// 参考記事「Goとクリーンアーキテクチャとトランザクションと」
// https://qiita.com/miya-masa/items/316256924a1f0d7374bb
type Transaction interface {
	DoInTx(context.Context, func(context.Context) error) error
}
