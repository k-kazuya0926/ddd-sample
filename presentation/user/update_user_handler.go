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
	UserID string
	Name   string
}

type UpdateUserResponse struct {
}

func (h *UpdateUserHandler) Handle(request UpdateUserRequest) UpdateUserResponse {
	_, err := h.useCase.Execute(user.UpdateUserUseCaseInput{
		ID:   request.UserID,
		Name: request.Name,
	})
	if err != nil {
		panic(err)
	}

	return UpdateUserResponse{}
}
