package task

import (
	"ddd-sample/presentation/shared/error"
	"ddd-sample/usecase/task"
	"time"
)

type FetchTaskDetailHandler struct {
	useCase task.FetchTaskDetailUseCase
}

func NewFetchTaskDetailHandler(useCase task.FetchTaskDetailUseCase) *FetchTaskDetailHandler {
	return &FetchTaskDetailHandler{useCase: useCase}
}

type FetchTaskDetailRequest struct {
	TaskID string
}

type FetchTaskDetailResponse struct {
	TaskID        string
	TaskName      string
	UserName      string
	TaskStatus    string
	PostponeCount uint64
	DueDate       time.Time
}

func (h *FetchTaskDetailHandler) Handle(request FetchTaskDetailRequest) FetchTaskDetailResponse {
	dto, err := h.useCase.Execute(task.FetchTaskDetailParam{
		TaskID: request.TaskID,
	})
	if err != nil {
		error.HandleError(err)
		return FetchTaskDetailResponse{}
	}

	return FetchTaskDetailResponse{
		TaskID:        dto.TaskID,
		TaskName:      dto.TaskName,
		UserName:      dto.UserName,
		TaskStatus:    dto.TaskStatus,
		PostponeCount: dto.PostponeCount,
		DueDate:       dto.DueDate,
	}
}
