package main

import "ddd-sample/presentation/user"

type Registry struct {
	createUserHandler *user.CreateUserHandler
	fetchUserHandler  *user.FetchUserHandler
	updateUserHandler *user.UpdateUserHandler
}

func NewRegistry(
	createUserHandler *user.CreateUserHandler,
	fetchUserHandler *user.FetchUserHandler,
	updateUserHandler *user.UpdateUserHandler,
) *Registry {
	return &Registry{
		createUserHandler: createUserHandler,
		fetchUserHandler:  fetchUserHandler,
		updateUserHandler: updateUserHandler,
	}
}
