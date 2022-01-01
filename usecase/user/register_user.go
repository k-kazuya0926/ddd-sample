package user

import (
	"context"
	domain "ddd-sample/domain/user"
	"ddd-sample/usecase/transaction"
	"errors"
)

type RegisterUserUseCase interface {
	Execute(input RegisterUserUseCaseInput) (RegisterUserUseCaseDTO, error)
}

type registerUserUseCase struct {
	transaction    transaction.Transaction
	userFactory    domain.UserFactory
	userRepository domain.UserRepository
}

func NewRegisterUserUseCase(
	transaction transaction.Transaction,
	userFactory domain.UserFactory,
	userRepository domain.UserRepository,
) RegisterUserUseCase {
	return &registerUserUseCase{
		transaction:    transaction,
		userFactory:    userFactory,
		userRepository: userRepository,
	}
}

type RegisterUserUseCaseInput struct {
	Name string
}

type RegisterUserUseCaseDTO struct {
	UserID string
}

func (uc *registerUserUseCase) Execute(input RegisterUserUseCaseInput) (RegisterUserUseCaseDTO, error) {
	var user domain.User
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		userName, err := domain.NewUserName(input.Name)
		if err != nil {
			return err
		}
		user = uc.userFactory.Create(userName)

		// 重複チェック
		userDuplicationChecker := domain.NewUserDuplicationChecker(uc.userRepository)
		userExists, err := userDuplicationChecker.Exists(ctx, user)
		if err != nil {
			return err
		}
		if userExists {
			return errors.New("すでに登録されています。")
		}

		err = uc.userRepository.Insert(ctx, user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return RegisterUserUseCaseDTO{}, err
	}
	return RegisterUserUseCaseDTO{UserID: user.ID().String()}, nil
}
