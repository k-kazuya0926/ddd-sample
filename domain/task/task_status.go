package task

type TaskStatus uint64

const (
	TaskStatusUnDone TaskStatus = iota + 1
	TaskStatusDone
)

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusUnDone:
		return "未完了"
	case TaskStatusDone:
		return "完了"
	default:
		return ""
	}
}
