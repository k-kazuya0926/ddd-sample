//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package user

import "context"

type UserRepository interface {
	Insert(ctx context.Context, user User) error
	FindByName(ctx context.Context, name UserName) (*User, error)
	FindByID(ctx context.Context, id UserID) (*User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id UserID) error
}
