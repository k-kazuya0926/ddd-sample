package task

import (
	"context"
	domain "ddd-sample/domain/task"
	usecase_error "ddd-sample/usecase/shared/error"
	"ddd-sample/usecase/shared/transaction"
	"fmt"
	"time"
)

type FetchTaskDetailUseCase interface {
	Execute(input FetchTaskDetailUseCaseInput) (FetchTaskDetailUseCaseDTO, error)
}

type fetchTaskDetailUseCase struct {
	transaction    transaction.Transaction
	taskRepository domain.TaskRepository
}

func NewFetchTaskDetailUseCase(
	transaction transaction.Transaction,
	taskRepository domain.TaskRepository,
) FetchTaskDetailUseCase {
	return &fetchTaskDetailUseCase{
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

type FetchTaskDetailUseCaseInput struct {
	TaskID string
}

type FetchTaskDetailUseCaseDTO struct {
	TaskID        string
	TaskName      string
	TaskStatus    string
	PostponeCount uint64
	DueDate       time.Time
}

func (uc *fetchTaskDetailUseCase) Execute(input FetchTaskDetailUseCaseInput) (FetchTaskDetailUseCaseDTO, error) {
	var task *domain.Task
	err := uc.transaction.DoInTx(context.Background(), func(ctx context.Context) error {
		taskID, err := domain.ParseTaskID(input.TaskID)
		if err != nil {
			return err
		}
		task, err = uc.taskRepository.FindByID(ctx, taskID)
		if err != nil {
			return err
		}
		if task == nil {
			return fmt.Errorf("%w", usecase_error.NewUseCaseError("タスクが存在しません。"))
		}

		return nil
	})
	if err != nil {
		return FetchTaskDetailUseCaseDTO{}, err
	}
	return FetchTaskDetailUseCaseDTO{
		TaskID:        task.ID().String(),
		TaskName:      task.Name().String(),
		TaskStatus:    task.Stasus().String(),
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}, nil
}
