package user

import (
	"ddd-sample/usecase/user"
)

type UpdateUserHandler struct {
	useCase user.UpdateUserUseCase
}

func NewUpdateUserHandler(useCase user.UpdateUserUseCase) *UpdateUserHandler {
	return &UpdateUserHandler{useCase: useCase}
}

type UpdateUserRequest struct {
	UserID   string
	UserName string
}

type UpdateUserResponse struct {
}

func (h *UpdateUserHandler) Handle(request UpdateUserRequest) UpdateUserResponse {
	_, err := h.useCase.Execute(user.UpdateUserUseCaseInput{
		UserID:   request.UserID,
		UserName: request.UserName,
	})
	if err != nil {
		panic(err)
	}

	return UpdateUserResponse{}
}
