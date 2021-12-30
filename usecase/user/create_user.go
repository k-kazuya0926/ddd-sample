package user

import (
	"context"
	domain "ddd-sample/domain/user"
	"ddd-sample/usecase/transaction"
	"errors"
)

type CreateUserUseCase interface {
	Execute(input CreateUserUseCaseInput) (CreateUserUseCaseDTO, error)
}

type createUserUseCase struct {
	userRepository domain.UserRepository
	transaction    transaction.Transaction
}

func NewCreateUserUseCase(userRepository domain.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}

type CreateUserUseCaseInput struct {
	Name string
}

type CreateUserUseCaseDTO struct {
}

func (uc *createUserUseCase) Execute(input CreateUserUseCaseInput) (CreateUserUseCaseDTO, error) {
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		userName, err := domain.NewUserName(input.Name)
		if err != nil {
			return err
		}
		user := domain.NewUser(userName)

		// 重複チェック
		userDuplicationChecker := domain.NewUserDuplicationChecker(uc.userRepository)
		userExists, err := userDuplicationChecker.Exists(user)
		if err != nil {
			return err
		}
		if userExists {
			return errors.New("すでに登録されています。")
		}

		err = uc.userRepository.Insert(user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return CreateUserUseCaseDTO{}, err
	}
	return CreateUserUseCaseDTO{}, nil
}
