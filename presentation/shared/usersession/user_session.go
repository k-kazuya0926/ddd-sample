package usersession

import (
	"ddd-sample/domain/user"
)

type userSession struct {
	userID user.UserID
}

func (s *userSession) UserID() user.UserID {
	return s.userID
}
