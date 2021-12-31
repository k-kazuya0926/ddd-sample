//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package user

type UserFactory interface {
	Create(name UserName) User
}

type userFactory struct {
}

func NewUserFactory() UserFactory {
	return &userFactory{}
}

func (uf *userFactory) Create(name UserName) User {
	return User{
		id:   newUserID(),
		name: name,
	}
}
