package user

import (
	"ddd-sample/presentation/shared/error"
	"ddd-sample/usecase/user"
)

type RegisterUserHandler struct {
	useCase user.RegisterUserUseCase
}

func NewRegisterUserHandler(useCase user.RegisterUserUseCase) *RegisterUserHandler {
	return &RegisterUserHandler{useCase: useCase}
}

type RegisterUserRequest struct {
	UserName string
}

type RegisterUserResponse struct {
	UserID string
}

func (h *RegisterUserHandler) Handle(request RegisterUserRequest) RegisterUserResponse {
	dto, err := h.useCase.Execute(user.RegisterUserParam{
		UserName: request.UserName,
	})
	if err != nil {
		error.HandleError(err)
		return RegisterUserResponse{}
	}

	return RegisterUserResponse{UserID: dto.UserID}
}
