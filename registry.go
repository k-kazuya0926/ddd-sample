package main

import "ddd-sample/presentation/user"

type Registry struct {
	createUserHandler *user.CreateUserHandler
	updateUserHandler *user.UpdateUserHandler
}

func NewRegistry(
	createUserHandler *user.CreateUserHandler,
	updateUserHandler *user.UpdateUserHandler,
) *Registry {
	return &Registry{
		createUserHandler: createUserHandler,
		updateUserHandler: updateUserHandler,
	}
}
