package task

import (
	"fmt"
	"unicode/utf8"
)

type TaskName struct {
	name string
}

const taskNameMaxLength = 10

func NewTaskName(name string) (TaskName, error) {
	if utf8.RuneCountInString(name) > taskNameMaxLength {
		return TaskName{}, fmt.Errorf("タスク名は%d文字以下で入力してください", taskNameMaxLength)
	}
	return TaskName{name: name}, nil
}
