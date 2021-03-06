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
		UserName: "ユーザー1",
	})
	fmt.Printf("registerUserResponse: %+v\n", registerUserResponse)

	// ユーザー詳細取得
	fetchUserDetailResponse := registry.fetchUserDetailHandler.Handle(user.FetchUserDetailRequest{
		UserID: registerUserResponse.UserID,
	})
	fmt.Printf("fetchUserDetailResponse: %+v\n", fetchUserDetailResponse)

	// ユーザー更新
	updateUserResponse := registry.updateUserHandler.Handle(user.UpdateUserRequest{
		UserID:   registerUserResponse.UserID,
		UserName: "ユーザー2",
	})
	fmt.Printf("updateUserResponse: %+v\n", updateUserResponse)

	// タスク作成
	createTaskResponse := registry.createTaskHandler.Handle(task.CreateTaskRequest{
		TaskName: "タスク1",
		DueDate:  time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
		UserID:   registerUserResponse.UserID,
	})
	fmt.Printf("createTaskResponse: %+v\n", createTaskResponse)

	// タスク詳細取得
	fetchTaskDetailResponse := registry.fetchTaskDetailHandler.Handle(task.FetchTaskDetailRequest{
		TaskID: createTaskResponse.TaskID,
	})
	fmt.Printf("fetchTaskDetailResponse: %+v\n", fetchTaskDetailResponse)

	// タスク延期
	postponeTaskResponse := registry.postponeTaskHandler.Handle(task.PostponeTaskRequest{
		TaskID: createTaskResponse.TaskID,
	})
	fmt.Printf("postponeTaskResponse: %+v\n", postponeTaskResponse)

	// ユーザー削除
	deleteUserResponse := registry.deleteUserHandler.Handle(user.DeleteUserRequest{
		UserID: registerUserResponse.UserID,
	})
	fmt.Printf("deleteUserResponse: %+v\n", deleteUserResponse)
}
