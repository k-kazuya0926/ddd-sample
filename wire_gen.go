// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-sample/domain/task"
	"ddd-sample/domain/user"
	task2 "ddd-sample/infra/in_memory/task"
	"ddd-sample/infra/in_memory/transaction"
	user2 "ddd-sample/infra/in_memory/user"
	task4 "ddd-sample/presentation/task"
	user4 "ddd-sample/presentation/user"
	task3 "ddd-sample/usecase/task"
	user3 "ddd-sample/usecase/user"
)

// Injectors from wire.go:

func initRegistry() *Registry {
	transactionTransaction := transaction.NewNoopTransaction()
	userFactory := user.NewUserFactory()
	userRepository := user2.NewInMemoryUserRepository()
	userDuplicationChecker := user.NewUserDuplicationChecker(userRepository)
	registerUserUseCase := user3.NewRegisterUserUseCase(transactionTransaction, userFactory, userDuplicationChecker, userRepository)
	registerUserHandler := user4.NewRegisterUserHandler(registerUserUseCase)
	fetchUserUseCase := user3.NewFetchUserUseCase(userRepository)
	fetchUserHandler := user4.NewFetchUserHandler(fetchUserUseCase)
	updateUserUseCase := user3.NewUpdateUserUseCase(transactionTransaction, userDuplicationChecker, userRepository)
	updateUserHandler := user4.NewUpdateUserHandler(updateUserUseCase)
	deleteUserUseCase := user3.NewDeleteUserUseCase(userRepository)
	deleteUserHandler := user4.NewDeleteUserHandler(deleteUserUseCase)
	taskFactory := task.NewTaskFactory()
	taskRepository := task2.NewInMemoryTaskRepository()
	createTaskUseCase := task3.NewCreateTaskUseCase(transactionTransaction, taskFactory, taskRepository)
	createTaskHandler := task4.NewCreateTaskHandler(createTaskUseCase)
	registry := NewRegistry(registerUserHandler, fetchUserHandler, updateUserHandler, deleteUserHandler, createTaskHandler)
	return registry
}
