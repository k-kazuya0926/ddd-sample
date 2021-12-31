package main

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/infra/transaction"
	infra "ddd-sample/infra/user"
	usecase "ddd-sample/usecase/user"
	"fmt"
)

func main() {
	// TODO DI
	transaction := transaction.NewNoopTransaction()
	userFactory := domain.NewUserFactory()
	userRepository := infra.NewInMemoryUserRepository()

	// ユーザー登録
	createUserUseCase := usecase.NewCreateUserUseCase(transaction, userFactory, userRepository)
	createUserUseCaseDTO, err := createUserUseCase.Execute(usecase.CreateUserUseCaseInput{Name: "ダミーユーザー"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("createUserUseCaseDTO: %+v\n", createUserUseCaseDTO)

	// ユーザー更新
	updateUserUseCase := usecase.NewUpdateUserUseCase(transaction, userRepository)
	updateUserUseCaseDTO, err := updateUserUseCase.Execute(usecase.UpdateUserUseCaseInput{
		ID:   createUserUseCaseDTO.UserID,
		Name: "ダミーユーザー2",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("updateUserUseCaseDTO: %+v\n", updateUserUseCaseDTO)
}
