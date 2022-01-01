// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-sample/domain/user"
	"ddd-sample/infra/in_memory/transaction"
	user2 "ddd-sample/infra/in_memory/user"
	user4 "ddd-sample/presentation/user"
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
	registry := NewRegistry(registerUserHandler, fetchUserHandler, updateUserHandler, deleteUserHandler)
	return registry
}
