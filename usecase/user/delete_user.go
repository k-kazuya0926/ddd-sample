package user

import (
	"context"
	domain "ddd-sample/domain/user"
)

type DeleteUserUseCase interface {
	Execute(input DeleteUserUseCaseInput) (DeleteUserUseCaseDTO, error)
}

type deleteUserUseCase struct {
	userRepository domain.UserRepository
}

func NewDeleteUserUseCase(userRepository domain.UserRepository) DeleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepository,
	}
}

type DeleteUserUseCaseInput struct {
	ID string
}

type DeleteUserUseCaseDTO struct {
}

func (uc *deleteUserUseCase) Execute(input DeleteUserUseCaseInput) (DeleteUserUseCaseDTO, error) {
	userID, err := domain.ParseUserID(input.ID)
	if err != nil {
		return DeleteUserUseCaseDTO{}, err
	}
	err = uc.userRepository.Delete(context.Background(), userID)
	if err != nil {
		return DeleteUserUseCaseDTO{}, err
	}

	return DeleteUserUseCaseDTO{}, nil
}
