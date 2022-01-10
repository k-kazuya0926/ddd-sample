package user

import (
	"context"
	domain "ddd-sample/domain/user"
)

type DeleteUserUseCase interface {
	Execute(input DeleteUserParam) (DeleteUserDTO, error)
}

type deleteUserUseCase struct {
	userRepository domain.UserRepository
}

func NewDeleteUserUseCase(userRepository domain.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

type DeleteUserParam struct {
	UserID string
}

type DeleteUserDTO struct {
}

func (uc *deleteUserUseCase) Execute(input DeleteUserParam) (DeleteUserDTO, error) {
	userID, err := domain.ParseUserID(input.UserID)
	if err != nil {
		return DeleteUserDTO{}, err
	}
	err = uc.userRepository.Delete(context.Background(), userID)
	if err != nil {
		return DeleteUserDTO{}, err
	}

	return DeleteUserDTO{}, nil
}
