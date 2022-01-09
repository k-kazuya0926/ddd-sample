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

func (r *InMemoryUserRepository) FindByName(ctx context.Context, userName domain.UserName) (*domain.User, error) {
	for _, user := range r.Store {
		if user.Name() == userName {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	for userID, user := range r.Store {
		if userID == userID {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user domain.User) error {
	r.Store[user.ID()] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, userID domain.UserID) error {
	delete(r.Store, userID)
	return nil
}
