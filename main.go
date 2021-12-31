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

	// ユーザー更新
	updateUserResponse := registry.updateUserHandler.Handle(user.UpdateUserRequest{
		UserID: createUserResponse.UserID,
		Name:   "ダミーユーザー2",
	})
	fmt.Printf("updateUserResponse: %+v\n", updateUserResponse)
}
