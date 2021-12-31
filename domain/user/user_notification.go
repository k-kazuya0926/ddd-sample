package user

// リポジトリにドメインモデルの内容を通知するためのインターフェース
type UserNotification interface {
	ID(id UserID)
	Name(name UserName)
}
