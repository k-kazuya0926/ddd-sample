package task

import (
	"context"
	domain "ddd-sample/domain/task"
	usecase_error "ddd-sample/usecase/shared/error"
	"ddd-sample/usecase/shared/transaction"
	"fmt"
	"time"
)

type PostponeTaskUseCase interface {
	Execute(input PostponeTaskParam) (PostponeTaskDTO, error)
}

type postponeTaskUseCase struct {
	transaction    transaction.Transaction
	taskRepository domain.TaskRepository
}

func NewPostponeTaskUseCase(
	transaction transaction.Transaction,
	taskRepository domain.TaskRepository,
) PostponeTaskUseCase {
	return &postponeTaskUseCase{
		transaction:    transaction,
		taskRepository: taskRepository,
	}
}

type PostponeTaskParam struct {
	TaskID string
}

type PostponeTaskDTO struct {
	PostponeCount uint64
	DueDate       time.Time
}

func (uc *postponeTaskUseCase) Execute(input PostponeTaskParam) (PostponeTaskDTO, error) {
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
			return fmt.Errorf("%w", usecase_error.TaskNotFoundError)
		}

		if err = task.Postpone(); err != nil {
			return err
		}

		err = uc.taskRepository.Update(ctx, *task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return PostponeTaskDTO{}, err
	}
	return PostponeTaskDTO{
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}, nil
}
