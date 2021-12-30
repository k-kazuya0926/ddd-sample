package user

import (
	"context"
	"errors"

	domain "ddd-sample/domain/user"
	"ddd-sample/usecase/transaction"
)

type UpdateUserUseCase interface {
	Execute(input UpdateUserUseCaseInput) (UpdateUserUseCaseDTO, error)
}

type updateUserUseCase struct {
	userRepository domain.UserRepository
	transaction    transaction.Transaction
}

type UpdateUserUseCaseInput struct {
	ID   string
	Name string
}

type UpdateUserUseCaseDTO struct {
}

func (uc *updateUserUseCase) Execute(input UpdateUserUseCaseInput) (UpdateUserUseCaseDTO, error) {
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		user, err := uc.userRepository.FindByID(input.ID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("User is not found.")
		}

		user.SetName(input.Name)

		// 重複チェック
		userDuplicationChecker := domain.NewUserDuplicationChecker(uc.userRepository)
		userExists, err := userDuplicationChecker.Exists(*user)
		if err != nil {
			return err
		}
		if userExists {
			return errors.New("Duplicate user exists.")
		}

		err = uc.userRepository.Update(*user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return UpdateUserUseCaseDTO{}, err
	}
	return UpdateUserUseCaseDTO{}, nil
}

func NewUpdateUserUseCase(userRepository domain.UserRepository) UpdateUserUseCase {
	return &updateUserUseCase{
		userRepository: userRepository,
	}
}
