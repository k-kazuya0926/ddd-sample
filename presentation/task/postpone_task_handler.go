package task

import (
	"ddd-sample/presentation/shared/error"
	"ddd-sample/usecase/task"
	"time"
)

type PostponeTaskHandler struct {
	useCase task.PostponeTaskUseCase
}

func NewPostponeTaskHandler(useCase task.PostponeTaskUseCase) *PostponeTaskHandler {
	return &PostponeTaskHandler{useCase: useCase}
}

type PostponeTaskRequest struct {
	TaskID string
}

type PostponeTaskResponse struct {
	PostponeCount uint64
	DueDate       time.Time
}

func (h *PostponeTaskHandler) Handle(request PostponeTaskRequest) PostponeTaskResponse {
	dto, err := h.useCase.Execute(task.PostponeTaskParam{
		TaskID: request.TaskID,
	})
	if err != nil {
		error.HandleError(err)
		return PostponeTaskResponse{}
	}

	return PostponeTaskResponse{
		PostponeCount: dto.PostponeCount,
		DueDate:       dto.DueDate,
	}
}
