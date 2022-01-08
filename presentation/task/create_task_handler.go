package task

import (
	"ddd-sample/usecase/task"
)

type CreateTaskHandler struct {
	useCase task.CreateTaskUseCase
}

func NewCreateTaskHandler(useCase task.CreateTaskUseCase) *CreateTaskHandler {
	return &CreateTaskHandler{useCase: useCase}
}

type CreateTaskRequest struct {
	Name string
}

type CreateTaskResponse struct {
	TaskID string
}

func (h *CreateTaskHandler) Handle(request CreateTaskRequest) CreateTaskResponse {
	dto, err := h.useCase.Execute(task.CreateTaskUseCaseInput{
		Name: request.Name,
	})
	if err != nil {
		panic(err)
	}

	return CreateTaskResponse{TaskID: dto.TaskID}
}
