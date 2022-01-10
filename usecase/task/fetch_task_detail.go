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
	Execute(input FetchTaskDetailParam) (FetchTaskDetailDTO, error)
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

type FetchTaskDetailParam struct {
	TaskID string
}

type FetchTaskDetailDTO struct {
	TaskID        string
	TaskName      string
	UserName      string
	TaskStatus    string
	PostponeCount uint64
	DueDate       time.Time
}

func (uc *fetchTaskDetailUseCase) Execute(input FetchTaskDetailParam) (FetchTaskDetailDTO, error) {
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
			return fmt.Errorf("%w", usecase_error.TaskNotFoundError)
		}

		user, err = uc.userRepository.FindByID(ctx, task.UserID())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return FetchTaskDetailDTO{}, err
	}
	return FetchTaskDetailDTO{
		TaskID:        task.ID().String(),
		TaskName:      task.Name().String(),
		UserName:      user.Name().String(),
		TaskStatus:    task.Stasus().String(),
		PostponeCount: task.PostponeCount(),
		DueDate:       task.DueDate(),
	}, nil
}
