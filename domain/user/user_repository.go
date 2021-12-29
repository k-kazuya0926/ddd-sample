//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package user

type UserRepository interface {
	Insert(user User) error
	Find(id string) (User, error)
}
