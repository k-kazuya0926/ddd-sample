package user

import (
	"ddd-sample/usecase/user"
)

type RegisterUserHandler struct {
	useCase user.RegisterUserUseCase
}

func NewRegisterUserHandler(useCase user.RegisterUserUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{useCase: useCase}
}

type RegisterUserRequest struct {
	Name string
}

type RegisterUserResponse struct {
	UserID string
}

func (h *RegisterUserHandler) Handle(request RegisterUserRequest) RegisterUserResponse {
	dto, err := h.useCase.Execute(user.RegisterUserUseCaseInput{
		Name: request.Name,
	})
	if err != nil {
		panic(err)
	}

	return RegisterUserResponse{UserID: dto.UserID}
}
