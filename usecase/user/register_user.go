package user

import (
	"context"
	domain "ddd-sample/domain/user"
	usecase_error "ddd-sample/usecase/shared/error"
	"ddd-sample/usecase/transaction"
	"fmt"
)

type RegisterUserUseCase interface {
	Execute(input RegisterUserUseCaseInput) (RegisterUserUseCaseDTO, error)
}

type registerUserUseCase struct {
	transaction            transaction.Transaction
	userFactory            domain.UserFactory
	userDuplicationChecker domain.UserDuplicationChecker
	userRepository         domain.UserRepository
}

func NewRegisterUserUseCase(
	transaction transaction.Transaction,
	userFactory domain.UserFactory,
	userDuplicationChecker domain.UserDuplicationChecker,
	userRepository domain.UserRepository,
) RegisterUserUseCase {
	return &registerUserUseCase{
		transaction:            transaction,
		userFactory:            userFactory,
		userDuplicationChecker: userDuplicationChecker,
		userRepository:         userRepository,
	}
}

type RegisterUserUseCaseInput struct {
	UserName string
}

type RegisterUserUseCaseDTO struct {
	UserID string
}

func (uc *registerUserUseCase) Execute(input RegisterUserUseCaseInput) (RegisterUserUseCaseDTO, error) {
	var user domain.User
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		userName, err := domain.NewUserName(input.UserName)
		if err != nil {
			return err
		}
		user = uc.userFactory.Create(userName)

		// 重複チェック
		userExists, err := uc.userDuplicationChecker.Exists(ctx, user)
		if err != nil {
			return err
		}
		if userExists {
			return fmt.Errorf("%w", usecase_error.NewUseCaseError("すでに登録されています。"))
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
