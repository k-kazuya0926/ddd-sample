package user

import (
	"ddd-sample/presentation/shared/error"
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
	UserID   string
	UserName string
}

func (h *FetchUserDetailHandler) Handle(request FetchUserDetailRequest) FetchUserDetailResponse {
	dto, err := h.useCase.Execute(user.FetchUserDetailUseCaseInput{
		UserID: request.UserID,
	})
	if err != nil {
		error.HandleError(err)
		return FetchUserDetailResponse{}
	}

	return FetchUserDetailResponse{
		UserID:   dto.UserID,
		UserName: dto.UserName,
	}
}
