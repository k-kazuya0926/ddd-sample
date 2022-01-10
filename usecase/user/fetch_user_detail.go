package user

import (
	"context"
	domain "ddd-sample/domain/user"
	usecase_error "ddd-sample/usecase/shared/error"
	"fmt"
)

type FetchUserDetailUseCase interface {
	Execute(input FetchUserDetailParam) (FetchUserDetailDTO, error)
}

type fetchUserDetailUseCase struct {
	userRepository domain.UserRepository
}

func NewFetchUserDetailUseCase(userRepository domain.UserRepository) FetchUserDetailUseCase {
	return &fetchUserDetailUseCase{
		userRepository: userRepository,
	}
}

type FetchUserDetailParam struct {
	UserID string
}

type FetchUserDetailDTO struct {
	UserID   string
	UserName string
}

func (uc *fetchUserDetailUseCase) Execute(input FetchUserDetailParam) (FetchUserDetailDTO, error) {
	userID, err := domain.ParseUserID(input.UserID)
	if err != nil {
		return FetchUserDetailDTO{}, err
	}
	user, err := uc.userRepository.FindByID(context.Background(), userID)
	if err != nil {
		return FetchUserDetailDTO{}, err
	}
	if user == nil {
		return FetchUserDetailDTO{}, fmt.Errorf("%w", usecase_error.UserNotFoundError)
	}

	return FetchUserDetailDTO{
		UserID:   user.ID().String(),
		UserName: user.Name().String(),
	}, nil
}
