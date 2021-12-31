package user

import (
	"ddd-sample/usecase/user"
)

type DeleteUserHandler struct {
	useCase user.DeleteUserUseCase
}

func NewDeleteUserHandler(useCase user.DeleteUserUseCase) *DeleteUserHandler {
	return &DeleteUserHandler{useCase: useCase}
}

type DeleteUserRequest struct {
	UserID string
}

type DeleteUserResponse struct {
}

func (h *DeleteUserHandler) Handle(request DeleteUserRequest) DeleteUserResponse {
	_, err := h.useCase.Execute(user.DeleteUserUseCaseInput{
		ID: request.UserID,
	})
	if err != nil {
		panic(err)
	}

	return DeleteUserResponse{}
}
