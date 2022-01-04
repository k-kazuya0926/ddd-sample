package task

type TaskName struct {
	name string
}

func NewTaskName(name string) TaskName {
	return TaskName{name: name}
}
