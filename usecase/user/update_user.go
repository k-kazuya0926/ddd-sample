package user

import (
	"context"
	domain "ddd-sample/domain/user"
	"ddd-sample/usecase/transaction"
	"errors"
)

type UpdateUserUseCase interface {
	Execute(input UpdateUserUseCaseInput) (UpdateUserUseCaseDTO, error)
}

type updateUserUseCase struct {
	transaction    transaction.Transaction
	userRepository domain.UserRepository
}

func NewUpdateUserUseCase(userRepository domain.UserRepository) UpdateUserUseCase {
	return &updateUserUseCase{
		userRepository: userRepository,
	}
}

type UpdateUserUseCaseInput struct {
	ID   string
	Name string
}

type UpdateUserUseCaseDTO struct {
}

func (uc *updateUserUseCase) Execute(input UpdateUserUseCaseInput) (UpdateUserUseCaseDTO, error) {
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		userID, err := domain.ParseUserID(input.ID)
		if err != nil {
			return err
		}
		user, err := uc.userRepository.FindByID(ctx, userID)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("ユーザーが存在しません。")
		}

		userName, err := domain.NewUserName(input.Name)
		if err != nil {
			return err
		}
		user.SetName(userName)

		// 重複チェック
		userDuplicationChecker := domain.NewUserDuplicationChecker(uc.userRepository)
		userExists, err := userDuplicationChecker.Exists(ctx, *user)
		if err != nil {
			return err
		}
		if userExists {
			return errors.New("すでに登録されています。")
		}

		err = uc.userRepository.Update(ctx, *user)
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
