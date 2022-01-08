package user

import (
	"context"
	domain "ddd-sample/domain/user"
	"errors"
)

type FetchUserDetailUseCase interface {
	Execute(input FetchUserDetailUseCaseInput) (FetchUserDetailUseCaseDTO, error)
}

type fetchUserDetailUseCase struct {
	userRepository domain.UserRepository
}

func NewFetchUserDetailUseCase(userRepository domain.UserRepository) FetchUserDetailUseCase {
	return &fetchUserDetailUseCase{
		userRepository: userRepository,
	}
}

type FetchUserDetailUseCaseInput struct {
	ID string
}

type FetchUserDetailUseCaseDTO struct {
	ID   string
	Name string
}

func (uc *fetchUserDetailUseCase) Execute(input FetchUserDetailUseCaseInput) (FetchUserDetailUseCaseDTO, error) {
	userID, err := domain.ParseUserID(input.ID)
	if err != nil {
		return FetchUserDetailUseCaseDTO{}, err
	}
	user, err := uc.userRepository.FindByID(context.Background(), userID)
	if err != nil {
		return FetchUserDetailUseCaseDTO{}, err
	}
	if user == nil {
		return FetchUserDetailUseCaseDTO{}, errors.New("ユーザーが存在しません。")
	}

	return FetchUserDetailUseCaseDTO{
		ID:   user.ID().String(),
		Name: user.Name().String(),
	}, nil
}
