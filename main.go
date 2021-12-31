package main

import (
	"ddd-sample/presentation/user"
	"fmt"
)

func main() {
	registry := initRegistry()

	// ユーザー登録
	createUserResponse := registry.createUserHandler.Handle(user.CreateUserRequest{
		Name: "ダミーユーザー",
	})
	fmt.Printf("createUserResponse: %+v\n", createUserResponse)

	// ユーザー取得
	fetchUserResponse := registry.fetchUserHandler.Handle(user.FetchUserRequest{
		UserID: createUserResponse.UserID,
	})
	fmt.Printf("fetchUserResponse: %+v\n", fetchUserResponse)

	// ユーザー更新
	updateUserResponse := registry.updateUserHandler.Handle(user.UpdateUserRequest{
		UserID: createUserResponse.UserID,
		Name:   "ダミーユーザー2",
	})
	fmt.Printf("updateUserResponse: %+v\n", updateUserResponse)

	// ユーザー削除
	deleteUserResponse := registry.deleteUserHandler.Handle(user.DeleteUserRequest{
		UserID: createUserResponse.UserID,
	})
	fmt.Printf("deleteUserResponse: %+v\n", deleteUserResponse)
}
