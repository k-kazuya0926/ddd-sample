package usersession

import (
	"ddd-sample/domain/user"
	"ddd-sample/usecase/shared/usersession"
)

type UserSessionProvider interface {
	getUserSession() (usersession.UserSession, error)
}

type userSessionProvider struct {
}

func (p *userSessionProvider) getUserSession() (usersession.UserSession, error) {
	// TODO 認証ライブラリに応じた実装

	userID, err := user.ParseUserID("dummy")
	if err != nil {
		return nil, err
	}
	return &userSession{userID: userID}, nil
}
