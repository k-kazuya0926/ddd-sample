//go:build wireinject
// +build wireinject

package main

import (
	domain_task "ddd-sample/domain/task"
	domain_user "ddd-sample/domain/user"
	in_memory_task "ddd-sample/infra/in_memory/task"
	"ddd-sample/infra/in_memory/transaction"
	in_memory_user "ddd-sample/infra/in_memory/user"
	presentation_task "ddd-sample/presentation/task"
	presentation_user "ddd-sample/presentation/user"
	usecase_task "ddd-sample/usecase/task"
	usecase_user "ddd-sample/usecase/user"

	"github.com/google/wire"
)

func initRegistry() *Registry {
	wire.Build(
		transaction.NewNoopTransaction,
		in_memory_user.NewInMemoryUserRepository,
		domain_user.NewUserFactory,
		domain_user.NewUserDuplicationChecker,
		usecase_user.NewRegisterUserUseCase,
		presentation_user.NewRegisterUserHandler,
		usecase_user.NewFetchUserDetailUseCase,
		presentation_user.NewFetchUserDetailHandler,
		usecase_user.NewUpdateUserUseCase,
		presentation_user.NewUpdateUserHandler,
		usecase_user.NewDeleteUserUseCase,
		presentation_user.NewDeleteUserHandler,
		in_memory_task.NewInMemoryTaskRepository,
		domain_task.NewTaskFactory,
		usecase_task.NewCreateTaskUseCase,
		presentation_task.NewCreateTaskHandler,
		NewRegistry,
	)
	return nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}
