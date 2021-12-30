//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package user

type UserRepository interface {
	Insert(user User) error
	FindByName(name string) (*User, error)
	FindByID(id string) (*User, error)
	Update(user User) error
}
