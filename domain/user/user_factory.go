//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package user

type UserFactory interface {
	Create(userName UserName) User
}

type userFactory struct {
}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func (f *userFactory) Create(userName UserName) User {
	return User{
		id:   newUserID(),
		name: userName,
	}
}
