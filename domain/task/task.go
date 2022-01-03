package task

import (
	"ddd-sample/domain/user"
	"time"
)

type Task struct {
	id            TaskID
	name          TaskName
	userID        user.UserID
	status        TaskStatus
	postponeCount uint64
	dueDate       time.Time
}

func NewTask(name TaskName, userID user.UserID, dueDate time.Time) Task {
	return Task{
		id:            NewTaskID(),
		name:          name,
		userID:        userID,
		status:        TaskStatusUnDone,
		postponeCount: 0,
		dueDate:       dueDate,
	}
}
