package task

import (
	"context"
	domain "ddd-sample/domain/task"
	usecase_error "ddd-sample/usecase/shared/error"
	"ddd-sample/usecase/transaction"
	"fmt"
	"time"
)

type PostponeTaskUseCase interface {
	Execute(input PostponeTaskUseCaseInput) (PostponeTaskUseCaseDTO, error)
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

type PostponeTaskUseCaseInput struct {
	TaskID string
}

type PostponeTaskUseCaseDTO struct {
	PostponeCount uint64
	DueDate       time.Time
}

func (uc *postponeTaskUseCase) Execute(input PostponeTaskUseCaseInput) (PostponeTaskUseCaseDTO, error) {
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

		task.Postpone()

		err = uc.taskRepository.Update(ctx, *task)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return PostponeTaskUseCaseDTO{}, err
	}
	return PostponeTaskUseCaseDTO{
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}, nil
}
