package usersession

import "ddd-sample/domain/user"

// 認証が必要なユースケースの引数にこれを渡す。interfaceとすることによりモックしやすくなる
type UserSession interface {
	UserID() user.UserID
	// UserRole() user.UserRole
}
