package main

import (
	"ddd-sample/presentation/task"
	"ddd-sample/presentation/user"
)

type Registry struct {
	registerUserHandler *user.RegisterUserHandler
	fetchUserHandler    *user.FetchUserHandler
	updateUserHandler   *user.UpdateUserHandler
	deleteUserHandler   *user.DeleteUserHandler
	createTaskHandler   *task.CreateTaskHandler
}

func NewRegistry(
	registerUserHandler *user.RegisterUserHandler,
	fetchUserHandler *user.FetchUserHandler,
	updateUserHandler *user.UpdateUserHandler,
	deleteUserHandler *user.DeleteUserHandler,
	createTaskHandler *task.CreateTaskHandler,
) *Registry {
	return &Registry{
		registerUserHandler: registerUserHandler,
		fetchUserHandler:    fetchUserHandler,
		updateUserHandler:   updateUserHandler,
		deleteUserHandler:   deleteUserHandler,
		createTaskHandler:   createTaskHandler,
	}
}
