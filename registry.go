package main

import (
	"ddd-sample/presentation/task"
	"ddd-sample/presentation/user"
)

type Registry struct {
	registerUserHandler    *user.RegisterUserHandler
	fetchUserDetailHandler *user.FetchUserDetailHandler
	updateUserHandler      *user.UpdateUserHandler
	deleteUserHandler      *user.DeleteUserHandler
	createTaskHandler      *task.CreateTaskHandler
}

func NewRegistry(
	registerUserHandler *user.RegisterUserHandler,
	fetchUserDetailHandler *user.FetchUserDetailHandler,
	updateUserHandler *user.UpdateUserHandler,
	deleteUserHandler *user.DeleteUserHandler,
	createTaskHandler *task.CreateTaskHandler,
) *Registry {
	return &Registry{
		registerUserHandler:    registerUserHandler,
		fetchUserDetailHandler: fetchUserDetailHandler,
		updateUserHandler:      updateUserHandler,
		deleteUserHandler:      deleteUserHandler,
		createTaskHandler:      createTaskHandler,
	}
}
