package user

import "ddd-sample/domain/user"

// テーブルのレコードに相当する構造体
type UserDataModel struct {
	ID   user.UserID
	Name user.UserName
}
