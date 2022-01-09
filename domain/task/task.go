package task

import (
	domain_error "ddd-sample/domain/shared/error"
	"ddd-sample/domain/user"
	"fmt"
	"time"
)

const maxPostponeCount = 3

type Task struct {
	id            TaskID
	name          TaskName
	userID        user.UserID
	status        TaskStatus
	postponeCount uint64
	dueDate       time.Time
}

func ReconstructTask(
	id TaskID,
	name TaskName,
	userID user.UserID,
	status TaskStatus,
	postponeCount uint64,
	dueDate time.Time,
) Task {
	return Task{
		id:            id,
		name:          name,
		userID:        userID,
		status:        status,
		postponeCount: postponeCount,
		dueDate:       dueDate,
	}
}

func (t *Task) Postpone() error {
	if t.postponeCount >= maxPostponeCount {
		return fmt.Errorf("%w", domain_error.NewDomainError("最大延期回数を超えています"))
	}
	t.dueDate = t.dueDate.AddDate(0, 0, 1)
	t.postponeCount += 1
	return nil
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) Name() TaskName {
	return t.name
}

func (t *Task) Stasus() TaskStatus {
	return t.status
}

func (t *Task) PostponeCount() uint64 {
	return t.postponeCount
}

func (t *Task) DueDate() time.Time {
	return t.dueDate
}
