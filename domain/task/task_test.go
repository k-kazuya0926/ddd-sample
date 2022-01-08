package task

import (
	"ddd-sample/domain/user"
	"reflect"
	"testing"
	"time"
)

func TestReconstructTask(t *testing.T) {
	taskID := NewTaskID()
	taskName, _ := NewTaskName("タスク名")
	userID, _ := user.ParseUserID("12345678901234567890123456")
	status := TaskStatusDone
	var postponeCount uint64 = 2
	dueDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	type args struct {
		id            TaskID
		name          TaskName
		userID        user.UserID
		status        TaskStatus
		postponeCount uint64
		dueDate       time.Time
	}
	tests := []struct {
		name string
		args args
		want Task
	}{
		{
			name: "ReconstructTaskに値を渡すと、渡した値でインスタンスが生成される",
			args: args{
				id:            taskID,
				name:          taskName,
				userID:        userID,
				status:        status,
				postponeCount: postponeCount,
				dueDate:       dueDate,
			},
			want: Task{
				id:            taskID,
				name:          taskName,
				userID:        userID,
				status:        status,
				postponeCount: postponeCount,
				dueDate:       dueDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReconstructTask(tt.args.id, tt.args.name, tt.args.userID, tt.args.status, tt.args.postponeCount, tt.args.dueDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReconstructTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Postpone(t1 *testing.T) {
	taskID := NewTaskID()
	taskName, _ := NewTaskName("タスク名")
	userID, _ := user.ParseUserID("12345678901234567890123456")
	status := TaskStatusUnDone
	dueDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	type fields struct {
		postponeCount uint64
		dueDate       time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		wantErr  bool
		wantTask Task
	}{
		{
			name: "延期回数2→3回",
			fields: fields{
				postponeCount: 2,
				dueDate:       dueDate,
			},
			wantErr:  false,
			wantTask: ReconstructTask(taskID, taskName, userID, status, 3, dueDate.AddDate(0, 0, 1)),
		},
		{
			name: "延期回数3→4回",
			fields: fields{
				postponeCount: 3,
				dueDate:       dueDate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := ReconstructTask(taskID, taskName, userID, status, tt.fields.postponeCount, tt.fields.dueDate)
			if err := t.Postpone(); (err != nil) != tt.wantErr {
				t1.Errorf("Postpone() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !reflect.DeepEqual(t, tt.wantTask) {
				t1.Errorf("got = %v, want %v", t, tt.wantTask)
			}
		})
	}
}
