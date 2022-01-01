package main

import "ddd-sample/presentation/user"

type Registry struct {
	registerUserHandler *user.RegisterUserHandler
	fetchUserHandler    *user.FetchUserHandler
	updateUserHandler   *user.UpdateUserHandler
	deleteUserHandler   *user.DeleteUserHandler
}

func NewRegistry(
	registerUserHandler *user.RegisterUserHandler,
	fetchUserHandler *user.FetchUserHandler,
	updateUserHandler *user.UpdateUserHandler,
	deleteUserHandler *user.DeleteUserHandler,
) *Registry {
	return &Registry{
		registerUserHandler: registerUserHandler,
		fetchUserHandler:    fetchUserHandler,
		updateUserHandler:   updateUserHandler,
		deleteUserHandler:   deleteUserHandler,
	}
}
