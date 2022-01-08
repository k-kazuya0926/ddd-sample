package user

import (
	"ddd-sample/usecase/user"
)

type FetchUserDetailHandler struct {
	useCase user.FetchUserDetailUseCase
}

func NewFetchUserDetailHandler(useCase user.FetchUserDetailUseCase) *FetchUserDetailHandler {
	return &FetchUserDetailHandler{useCase: useCase}
}

type FetchUserDetailRequest struct {
	UserID string
}

type FetchUserDetailResponse struct {
	UserID string
	Name   string
}

func (h *FetchUserDetailHandler) Handle(request FetchUserDetailRequest) FetchUserDetailResponse {
	dto, err := h.useCase.Execute(user.FetchUserDetailUseCaseInput{
		ID: request.UserID,
	})
	if err != nil {
		panic(err)
	}

	return FetchUserDetailResponse{
		UserID: dto.ID,
		Name:   dto.Name,
	}
}
