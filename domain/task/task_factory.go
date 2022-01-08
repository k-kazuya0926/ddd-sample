//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package task

import (
	"ddd-sample/domain/user"
	"time"
)

type TaskFactory interface {
	Create(name TaskName, dueDate time.Time, userID user.UserID) Task
}

type taskFactory struct {
}

func NewTaskFactory() TaskFactory {
	return &taskFactory{}
}

func (f *taskFactory) Create(name TaskName, dueDate time.Time, userID user.UserID) Task {
	return Task{
		id:            newTaskID(),
		name:          name,
		userID:        userID,
		status:        TaskStatusUnDone,
		postponeCount: 0,
		dueDate:       dueDate,
	}
}
