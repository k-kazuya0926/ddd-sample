package user

import (
	"context"
	domain "ddd-sample/domain/user"
)

// テスト用のリポジトリ
type InMemoryUserRepository struct {
	Store map[domain.UserID]domain.User
}

func NewInMemoryUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{Store: make(map[domain.UserID]domain.User, 0)}
}

func (r *InMemoryUserRepository) Insert(ctx context.Context, user domain.User) error {
	r.Store[user.ID()] = user
	return nil
}

func (r *InMemoryUserRepository) FindByName(ctx context.Context, name domain.UserName) (*domain.User, error) {
	for _, user := range r.Store {
		if user.Name() == name {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	for userID, user := range r.Store {
		if userID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user domain.User) error {
	r.Store[user.ID()] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id domain.UserID) error {
	delete(r.Store, id)
	return nil
}
