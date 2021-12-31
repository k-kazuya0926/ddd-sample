package user

import (
	"ddd-sample/usecase/user"
)

type FetchUserHandler struct {
	useCase user.FetchUserUseCase
}

func NewFetchUserHandler(useCase user.FetchUserUseCase) *FetchUserHandler {
	return &FetchUserHandler{useCase: useCase}
}

type FetchUserRequest struct {
	UserID string
}

type FetchUserResponse struct {
	UserID string
	Name   string
}

func (h *FetchUserHandler) Handle(request FetchUserRequest) FetchUserResponse {
	dto, err := h.useCase.Execute(user.FetchUserUseCaseInput{
		ID: request.UserID,
	})
	if err != nil {
		panic(err)
	}

	return FetchUserResponse{
		UserID: dto.ID,
		Name:   dto.Name,
	}
}
