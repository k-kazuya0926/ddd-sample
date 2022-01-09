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
	fetchTaskDetailHandler *task.FetchTaskDetailHandler
	postponeTaskHandler    *task.PostponeTaskHandler
}

func NewRegistry(
	registerUserHandler *user.RegisterUserHandler,
	fetchUserDetailHandler *user.FetchUserDetailHandler,
	updateUserHandler *user.UpdateUserHandler,
	deleteUserHandler *user.DeleteUserHandler,
	createTaskHandler *task.CreateTaskHandler,
	fetchTaskDetailHandler *task.FetchTaskDetailHandler,
	postponeTaskHandler *task.PostponeTaskHandler,
) *Registry {
	return &Registry{
		registerUserHandler:    registerUserHandler,
		fetchUserDetailHandler: fetchUserDetailHandler,
		updateUserHandler:      updateUserHandler,
		deleteUserHandler:      deleteUserHandler,
		createTaskHandler:      createTaskHandler,
		fetchTaskDetailHandler: fetchTaskDetailHandler,
		postponeTaskHandler:    postponeTaskHandler,
	}
}
