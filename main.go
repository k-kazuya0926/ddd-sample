package main

import (
	"ddd-sample/presentation/user"
	"fmt"
)

func main() {
	registry := initRegistry()

	// ユーザー登録
	registerUserResponse := registry.registerUserHandler.Handle(user.RegisterUserRequest{
		Name: "ダミーユーザー",
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
		Name:   "ダミーユーザー2",
	})
	fmt.Printf("updateUserResponse: %+v\n", updateUserResponse)

	// ユーザー削除
	deleteUserResponse := registry.deleteUserHandler.Handle(user.DeleteUserRequest{
		UserID: registerUserResponse.UserID,
	})
	fmt.Printf("deleteUserResponse: %+v\n", deleteUserResponse)
}
