package main

import (
	"ddd-sample/presentation/task"
	"ddd-sample/presentation/user"
	"fmt"
	"time"
)

func main() {
	registry := initRegistry()

	// ユーザー登録
	registerUserResponse := registry.registerUserHandler.Handle(user.RegisterUserRequest{
		Name: "ユーザー1",
	})
	fmt.Printf("registerUserResponse: %+v\n", registerUserResponse)

	// ユーザー取得
	fetchUserResponse := registry.fetchUserHandler.Handle(user.FetchUserRequest{
		UserID: registerUserResponse.UserID,
	})
	fmt.Printf("fetchUserResponse: %+v\n", fetchUserResponse)

	// ユーザー更新
	updateUserResponse := registry.updateUserHandler.Handle(user.UpdateUserRequest{
		UserID: registerUserResponse.UserID,
		Name:   "ユーザー2",
	})
	fmt.Printf("updateUserResponse: %+v\n", updateUserResponse)

	// タスク作成
	createTaskResponse := registry.createTaskHandler.Handle(task.CreateTaskRequest{
		Name:    "タスク1",
		DueDate: time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
		UserID:  registerUserResponse.UserID,
	})
	fmt.Printf("createTaskResponse: %+v\n", createTaskResponse)

	// ユーザー削除
	deleteUserResponse := registry.deleteUserHandler.Handle(user.DeleteUserRequest{
		UserID: registerUserResponse.UserID,
	})
	fmt.Printf("deleteUserResponse: %+v\n", deleteUserResponse)
}
