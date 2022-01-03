package task

type TaskStatus uint64

const (
	TaskStatusUnDone TaskStatus = iota + 1
	TaskStatusDone
)
