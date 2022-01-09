package task

import (
	"ddd-sample/domain/user"
	"ddd-sample/usecase/task"
	"time"
)

type CreateTaskHandler struct {
	useCase task.CreateTaskUseCase
}

func NewCreateTaskHandler(useCase task.CreateTaskUseCase) *CreateTaskHandler {
	return &CreateTaskHandler{useCase: useCase}
}

type CreateTaskRequest struct {
	TaskName string
	DueDate  time.Time
	UserID   string
}

type CreateTaskResponse struct {
	TaskID string
}

func (h *CreateTaskHandler) Handle(request CreateTaskRequest) CreateTaskResponse {
	// ユースケースがタイプセーフになるので、値オブジェクトの生成をプレゼンテーション層で行うのははOK
	userID, err := user.ParseUserID(request.UserID)
	if err != nil {
		panic(err)
	}
	dto, err := h.useCase.Execute(task.CreateTaskUseCaseInput{
		TaskName: request.TaskName,
		DueDate:  request.DueDate,
		UserID:   userID,
	})
	if err != nil {
		panic(err)
	}

	return CreateTaskResponse{TaskID: dto.TaskID}
}
