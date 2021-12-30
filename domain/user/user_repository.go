//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package user

type UserRepository interface {
	// TODO contextを渡したほうがいいかも
	Insert(user User) error
	FindByName(name UserName) (*User, error)
	FindByID(id string) (*User, error)
	Update(user User) error
}
