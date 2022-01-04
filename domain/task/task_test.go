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
			got := NewTask(tt.args.name, tt.args.dueDate, tt.args.userID)

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

func TestTask_Postpone_success(t1 *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		want    Task
	}{
		{
			name:    "タスクを延期すると、期日が1日後になり延期回数が1回増える",
			wantErr: false,
			want: Task{
				postponeCount: 1,
				dueDate:       time.Date(2022, 1, 2, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			dueDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
			userID, _ := user.ParseUserID("12345678901234567890123456")
			t := NewTask(NewTaskName("タスク"), dueDate, userID)
			err := t.Postpone()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Postpone() error = %v, wantErr %v", err, tt.wantErr)
			}
			if t.dueDate != tt.want.dueDate {
				t1.Errorf("got = %v, want %v", t.dueDate, tt.want.dueDate)
			}
			if t.postponeCount != tt.want.postponeCount {
				t1.Errorf("got = %v, want %v", t.postponeCount, tt.want.postponeCount)
			}
		})
	}
}

func TestTask_Postpone_failure(t1 *testing.T) {
	tests := []struct {
		name           string
		wantErrMessage string
	}{
		{
			name:           "最大回数延期されている場合、再度延期すると例外が発生する",
			wantErrMessage: "最大延期回数を超えています",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			dueDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
			userID, _ := user.ParseUserID("12345678901234567890123456")
			t := NewTask(NewTaskName("タスク"), dueDate, userID)
			t.Postpone()
			t.Postpone()
			t.Postpone()
			if err := t.Postpone(); err.Error() != tt.wantErrMessage {
				t1.Errorf("Postpone() error = %v, wantErrMessage %v", err, tt.wantErrMessage)
			}
		})
	}
}
