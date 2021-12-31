// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ddd-sample/domain/user"
	"ddd-sample/infra/transaction"
	user2 "ddd-sample/infra/user"
	user4 "ddd-sample/presentation/user"
	user3 "ddd-sample/usecase/user"
)

// Injectors from wire.go:

func initRegistry() *Registry {
	transactionTransaction := transaction.NewNoopTransaction()
	userFactory := user.NewUserFactory()
	userRepository := user2.NewInMemoryUserRepository()
	createUserUseCase := user3.NewCreateUserUseCase(transactionTransaction, userFactory, userRepository)
	createUserHandler := user4.NewCreateUserHandler(createUserUseCase)
	fetchUserUseCase := user3.NewFetchUserUseCase(userRepository)
	fetchUserHandler := user4.NewFetchUserHandler(fetchUserUseCase)
	updateUserUseCase := user3.NewUpdateUserUseCase(transactionTransaction, userRepository)
	updateUserHandler := user4.NewUpdateUserHandler(updateUserUseCase)
	deleteUserUseCase := user3.NewDeleteUserUseCase(userRepository)
	deleteUserHandler := user4.NewDeleteUserHandler(deleteUserUseCase)
	registry := NewRegistry(createUserHandler, fetchUserHandler, updateUserHandler, deleteUserHandler)
	return registry
}
