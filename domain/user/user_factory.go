//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package user

// ID採番処理をコントロールするために導入
type UserFactory interface {
	Create(name UserName) User
}

type userFactory struct {
}

func (uf *userFactory) Create(name UserName) User {
	return User{
		id:   "", // TODO ULIDを使うよう変更
		name: name,
	}
}
