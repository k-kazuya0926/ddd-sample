package main

import "ddd-sample/presentation/user"

type Registry struct {
	createUserHandler *user.CreateUserHandler
	fetchUserHandler  *user.FetchUserHandler
	updateUserHandler *user.UpdateUserHandler
	deleteUserHandler *user.DeleteUserHandler
}

func NewRegistry(
	createUserHandler *user.CreateUserHandler,
	fetchUserHandler *user.FetchUserHandler,
	updateUserHandler *user.UpdateUserHandler,
	deleteUserHandler *user.DeleteUserHandler,
) *Registry {
	return &Registry{
		createUserHandler: createUserHandler,
		fetchUserHandler:  fetchUserHandler,
		updateUserHandler: updateUserHandler,
		deleteUserHandler: deleteUserHandler,
	}
}
