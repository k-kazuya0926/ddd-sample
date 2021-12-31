//go:build wireinject
// +build wireinject

package main

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/infra/transaction"
	infra "ddd-sample/infra/user"
	presentation "ddd-sample/presentation/user"
	usecase "ddd-sample/usecase/user"

	"github.com/google/wire"
)

func initRegistry() *Registry {
	wire.Build(
		transaction.NewNoopTransaction,
		domain.NewUserFactory,
		infra.NewInMemoryUserRepository,
		usecase.NewCreateUserUseCase,
		presentation.NewCreateUserHandler,
		usecase.NewUpdateUserUseCase,
		presentation.NewUpdateUserHandler,
		NewRegistry,
	)
	return nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}