package user

import (
	"ddd-sample/usecase/user"
)

type CreateUserHandler struct {
	useCase user.CreateUserUseCase
}

func NewCreateUserHandler(useCase user.CreateUserUseCase) *CreateUserHandler {
	return &CreateUserHandler{useCase: useCase}
}

type CreateUserRequest struct {
	Name string
}

type CreateUserResponse struct {
	UserID string
}

func (h *CreateUserHandler) Handle(request CreateUserRequest) CreateUserResponse {
	dto, err := h.useCase.Execute(user.CreateUserUseCaseInput{
		Name: request.Name,
	})
	if err != nil {
		panic(err)
	}

	return CreateUserResponse{UserID: dto.UserID}
}
