package user

import (
	"context"
	domain "ddd-sample/domain/user"
	"errors"
)

type FetchUserUseCase interface {
	Execute(input FetchUserUseCaseInput) (FetchUserUseCaseDTO, error)
}

type fetchUserUseCase struct {
	userRepository domain.UserRepository
}

func NewFetchUserUseCase(userRepository domain.UserRepository) FetchUserUseCase {
	return &fetchUserUseCase{
		userRepository: userRepository,
	}
}

type FetchUserUseCaseInput struct {
	ID string
}

type FetchUserUseCaseDTO struct {
	ID   string
	Name string
}

func (uc *fetchUserUseCase) Execute(input FetchUserUseCaseInput) (FetchUserUseCaseDTO, error) {
	userID, err := domain.ParseUserID(input.ID)
	if err != nil {
		return FetchUserUseCaseDTO{}, err
	}
	user, err := uc.userRepository.FindByID(context.Background(), userID)
	if err != nil {
		return FetchUserUseCaseDTO{}, err
	}
	if user == nil {
		return FetchUserUseCaseDTO{}, errors.New("ユーザーが存在しません。")
	}

	return FetchUserUseCaseDTO{
		ID:   user.ID().String(),
		Name: user.Name().String(),
	}, nil
}
