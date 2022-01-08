package task

import (
	"context"
	domain "ddd-sample/domain/task"
	"ddd-sample/domain/user"
	"ddd-sample/usecase/transaction"
	"time"
)

type CreateTaskUseCase interface {
	Execute(input CreateTaskUseCaseInput) (CreateTaskUseCaseDTO, error)
}

type createTaskUseCase struct {
	transaction    transaction.Transaction
	taskRepository domain.TaskRepository
}

func NewCreateTaskUseCase(transaction transaction.Transaction, taskRepository domain.TaskRepository) CreateTaskUseCase {
	return &createTaskUseCase{
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

type CreateTaskUseCaseInput struct {
	Name    string
	DueDate time.Time
	UserID  user.UserID
}

type CreateTaskUseCaseDTO struct {
	TaskID string
}

func (uc *createTaskUseCase) Execute(input CreateTaskUseCaseInput) (CreateTaskUseCaseDTO, error) {
	var task domain.Task
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		taskNameModel, err := domain.NewTaskName(input.Name)
		if err != nil {
			return err
		}
		task = domain.NewTask(taskNameModel, input.DueDate, input.UserID)
		if err = uc.taskRepository.Insert(ctx, task); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return CreateTaskUseCaseDTO{}, err
	}
	return CreateTaskUseCaseDTO{TaskID: task.ID().String()}, nil
}
