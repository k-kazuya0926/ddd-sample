package task

import (
	"context"
	domain "ddd-sample/domain/task"
	"ddd-sample/domain/user"
	"ddd-sample/usecase/shared/transaction"
	"time"
)

type CreateTaskUseCase interface {
	Execute(input CreateTaskParam) (CreateTaskDTO, error)
}

type createTaskUseCase struct {
	transaction    transaction.Transaction
	taskFactory    domain.TaskFactory
	taskRepository domain.TaskRepository
}

func NewCreateTaskUseCase(
	transaction transaction.Transaction,
	taskFactory domain.TaskFactory,
	taskRepository domain.TaskRepository,
) CreateTaskUseCase {
	return &createTaskUseCase{
		transaction:    transaction,
		taskFactory:    taskFactory,
		taskRepository: taskRepository,
	}
}

type CreateTaskParam struct {
	TaskName string
	DueDate  time.Time
	UserID   user.UserID
}

type CreateTaskDTO struct {
	TaskID string
}

func (uc *createTaskUseCase) Execute(input CreateTaskParam) (CreateTaskDTO, error) {
	var task domain.Task
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		taskName, err := domain.NewTaskName(input.TaskName)
		if err != nil {
			return err
		}
		task = uc.taskFactory.Create(taskName, input.DueDate, input.UserID)
		if err = uc.taskRepository.Insert(ctx, task); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return CreateTaskDTO{}, err
	}
	return CreateTaskDTO{TaskID: task.ID().String()}, nil
}
