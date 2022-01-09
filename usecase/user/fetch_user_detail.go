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
	UserID string
}

type FetchUserDetailUseCaseDTO struct {
	UserID   string
	UserName string
}

func (uc *fetchUserDetailUseCase) Execute(input FetchUserDetailUseCaseInput) (FetchUserDetailUseCaseDTO, error) {
	userID, err := domain.ParseUserID(input.UserID)
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
		UserID:   user.ID().String(),
		UserName: user.Name().String(),
	}, nil
}
