package task

import (
	"ddd-sample/domain/user"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	taskName := TaskName{name: "A"}
	userID, _ := user.ParseUserID("12345678901234567890123456")
	dueDate := time.Now()
	type args struct {
		name    TaskName
		userID  user.UserID
		dueDate time.Time
	}
	tests := []struct {
		name string
		args args
		want Task
	}{
		{
			name: "新しくタスクを作成すると、未完了で延期回数0のインスタンスが生成される",
			args: args{
				name:    taskName,
				userID:  userID,
				dueDate: dueDate,
			},
			want: Task{
				name:          taskName,
				userID:        userID,
				status:        TaskStatusUnDone,
				postponeCount: 0,
				dueDate:       dueDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTask(tt.args.name, tt.args.userID, tt.args.dueDate)
			if got.name != tt.want.name {
				t.Errorf("got = %v, want %v", got.name, tt.want.name)
			}
			if got.userID != tt.want.userID {
				t.Errorf("got = %v, want %v", got.userID, tt.want.userID)
			}
			if got.status != tt.want.status {
				t.Errorf("got = %v, want %v", got.status, tt.want.status)
			}
			if got.postponeCount != tt.want.postponeCount {
				t.Errorf("got = %v, want %v", got.postponeCount, tt.want.postponeCount)
			}
			if got.dueDate != tt.want.dueDate {
				t.Errorf("got = %v, want %v", got.dueDate, tt.want.dueDate)
			}
		})
	}
}
