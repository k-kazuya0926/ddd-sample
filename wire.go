//go:build wireinject
// +build wireinject

package main

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/infra/in_memory/transaction"
	infra "ddd-sample/infra/in_memory/user"
	presentation "ddd-sample/presentation/user"
	usecase "ddd-sample/usecase/user"

	"github.com/google/wire"
)

func initRegistry() *Registry {
	wire.Build(
		transaction.NewNoopTransaction,
		domain.NewUserFactory,
		infra.NewInMemoryUserRepository,
		domain.NewUserDuplicationChecker,
		usecase.NewRegisterUserUseCase,
		presentation.NewRegisterUserHandler,
		usecase.NewFetchUserUseCase,
		presentation.NewFetchUserHandler,
		usecase.NewUpdateUserUseCase,
		presentation.NewUpdateUserHandler,
		usecase.NewDeleteUserUseCase,
		presentation.NewDeleteUserHandler,
		NewRegistry,
	)
	return nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}
