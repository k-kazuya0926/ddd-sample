package user

import (
	"context"
	domain "ddd-sample/domain/user"
	usecase_error "ddd-sample/usecase/shared/error"
	"ddd-sample/usecase/shared/transaction"
	"fmt"
)

type UpdateUserUseCase interface {
	Execute(input UpdateUserParam) (UpdateUserDTO, error)
}

type updateUserUseCase struct {
	transaction            transaction.Transaction
	userDuplicationChecker domain.UserDuplicationChecker
	userRepository         domain.UserRepository
}

func NewUpdateUserUseCase(
	transaction transaction.Transaction,
	userDuplicationChecker domain.UserDuplicationChecker,
	userRepository domain.UserRepository,
) UpdateUserUseCase {
	return &updateUserUseCase{
		transaction:            transaction,
		userDuplicationChecker: userDuplicationChecker,
		userRepository:         userRepository,
	}
}

type UpdateUserParam struct {
	UserID   string
	UserName string
}

type UpdateUserDTO struct {
}

func (uc *updateUserUseCase) Execute(input UpdateUserParam) (UpdateUserDTO, error) {
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		userID, err := domain.ParseUserID(input.UserID)
		if err != nil {
			return err
		}
		user, err := uc.userRepository.FindByID(ctx, userID)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("%w", usecase_error.NewUseCaseError("ユーザーが存在しません。"))
		}

		userName, err := domain.NewUserName(input.UserName)
		if err != nil {
			return err
		}
		user.ChangeName(userName)

		// 重複チェック
		userExists, err := uc.userDuplicationChecker.Exists(ctx, *user)
		if err != nil {
			return err
		}
		if userExists {
			return fmt.Errorf("%w", usecase_error.NewUseCaseError("すでに登録されています。"))
		}

		err = uc.userRepository.Update(ctx, *user)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return UpdateUserDTO{}, err
	}
	return UpdateUserDTO{}, nil
}
