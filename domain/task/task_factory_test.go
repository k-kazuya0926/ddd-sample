package task

import (
	"ddd-sample/domain/user"
	"testing"
	"time"
)

func Test_taskFactory_Create(t *testing.T) {
	taskName := TaskName{name: "A"}
	userID, _ := user.ParseUserID("12345678901234567890123456")
	dueDate := time.Now()
	type args struct {
		name    TaskName
		dueDate time.Time
		userID  user.UserID
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
			f := &taskFactory{}
			got := f.Create(tt.args.name, tt.args.dueDate, tt.args.userID)

			// TaskIDは比較対象外

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
