package task

import (
	"context"
	domain "ddd-sample/domain/task"
	domain_user "ddd-sample/domain/user"
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
	userRepository domain_user.UserRepository
}

func NewFetchTaskDetailUseCase(
	transaction transaction.Transaction,
	taskRepository domain.TaskRepository,
	userRepository domain_user.UserRepository,
) FetchTaskDetailUseCase {
	return &fetchTaskDetailUseCase{
		transaction:    transaction,
		taskRepository: taskRepository,
		userRepository: userRepository,
	}
}

type FetchTaskDetailUseCaseInput struct {
	TaskID string
}

type FetchTaskDetailUseCaseDTO struct {
	TaskID        string
	TaskName      string
	UserName      string
	TaskStatus    string
	PostponeCount uint64
	DueDate       time.Time
}

func (uc *fetchTaskDetailUseCase) Execute(input FetchTaskDetailUseCaseInput) (FetchTaskDetailUseCaseDTO, error) {
	var task *domain.Task
	var user *domain_user.User
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

		user, err = uc.userRepository.FindByID(ctx, task.UserID())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return FetchTaskDetailUseCaseDTO{}, err
	}
	return FetchTaskDetailUseCaseDTO{
		TaskID:        task.ID().String(),
		TaskName:      task.Name().String(),
		UserName:      user.Name().String(),
		TaskStatus:    task.Stasus().String(),
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}, nil
}
