package task

import (
	"context"
	domain "ddd-sample/domain/task"
)

// テスト用のリポジトリ
type InMemoryTaskRepository struct {
	Store map[domain.TaskID]domain.Task
}

func NewInMemoryTaskRepository() domain.TaskRepository {
	return &InMemoryTaskRepository{Store: make(map[domain.TaskID]domain.Task, 0)}
}

func (r *InMemoryTaskRepository) Insert(ctx context.Context, task domain.Task) error {
	r.Store[task.ID()] = task
	return nil
}

func (r *InMemoryTaskRepository) FindByID(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	for taskID, task := range r.Store {
		if taskID == taskID {
			return &task, nil
		}
	}
	return nil, nil
}

func (r *InMemoryTaskRepository) Update(ctx context.Context, task domain.Task) error {
	r.Store[task.ID()] = task
	return nil
}
