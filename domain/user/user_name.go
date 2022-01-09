package user

import (
	domain_error "ddd-sample/domain/shared/error"
	"fmt"
	"unicode/utf8"
)

// 値オブジェクト
type UserName struct {
	name string
}

const minUserNameLength = 3

func NewUserName(name string) (UserName, error) {
	// ユーザー名の文字数制限はドメイン知識なのでここに記述
	if utf8.RuneCountInString(name) < minUserNameLength {
		return UserName{}, fmt.Errorf("%w", domain_error.NewDomainError(fmt.Sprintf("ユーザー名は%d文字以上で入力してください。", minUserNameLength)))
	}
	return UserName{name: name}, nil
}

func (n UserName) String() string {
	return n.name
}
