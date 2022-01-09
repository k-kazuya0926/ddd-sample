package task

import (
	domain_task "ddd-sample/domain/task"
	"ddd-sample/domain/task/mock_task"
	domain_user "ddd-sample/domain/user"
	"ddd-sample/infra/in_memory/transaction"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_fetchTaskDetailUseCase_Execute(t *testing.T) {
	var (
		dummyUserIDString   = "12345678901234567890123456"
		dummyUserID, _      = domain_user.ParseUserID(dummyUserIDString)
		dummyTaskIDString   = "12345678901234567890123456"
		dummyTaskID, _      = domain_task.ParseTaskID(dummyTaskIDString)
		dummyTaskNameString = "ダミータスク"
		dummyTaskName, _    = domain_task.NewTaskName(dummyTaskNameString)
		dummyDueDate        = time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local)
	)
	type args struct {
		input FetchTaskDetailUseCaseInput
	}
	tests := []struct {
		name          string
		prepareMockFn func(*mock_task.MockTaskRepository)
		args          args
		want          FetchTaskDetailUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockTaskRepository *mock_task.MockTaskRepository) {
				task := domain_task.ReconstructTask(
					dummyTaskID,
					dummyTaskName,
					dummyUserID,
					domain_task.TaskStatusUnDone,
					0,
					dummyDueDate,
				)
				mockTaskRepository.EXPECT().FindByID(gomock.Any(), dummyTaskID).Return(&task, nil)
			},
			args: args{
				input: FetchTaskDetailUseCaseInput{
					TaskID: dummyTaskIDString,
				},
			},
			want: FetchTaskDetailUseCaseDTO{
				TaskID:        dummyTaskIDString,
				TaskName:      dummyTaskNameString,
				TaskStatus:    "未完了",
				PostponeCount: 0,
				DueDate:       dummyDueDate,
			},
			wantErr: false,
		},
		{
			name: "異常系：タスクが存在しない",
			prepareMockFn: func(mockTaskRepository *mock_task.MockTaskRepository) {
				mockTaskRepository.EXPECT().FindByID(gomock.Any(), dummyTaskID).Return(nil, nil)
			},
			args: args{
				input: FetchTaskDetailUseCaseInput{
					TaskID: dummyTaskIDString,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockTaskRepository := mock_task.NewMockTaskRepository(mockCtrl)
			// 参考記事「gomockを完全に理解する」
			// https://zenn.dev/sanpo_shiho/articles/01da627ead98f5
			tt.prepareMockFn(mockTaskRepository)

			uc := &fetchTaskDetailUseCase{
				transaction:    transaction.NewNoopTransaction(),
				taskRepository: mockTaskRepository,
			}
			got, err := uc.Execute(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
