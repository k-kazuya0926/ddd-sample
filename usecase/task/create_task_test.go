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

func Test_createTaskUseCase_Execute(t *testing.T) {
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
		input CreateTaskUseCaseInput
	}
	tests := []struct {
		name          string
		prepareMockFn func(*mock_task.MockTaskFactory, *mock_task.MockTaskRepository)
		args          args
		want          CreateTaskUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockTaskFactory *mock_task.MockTaskFactory, mockTaskRepository *mock_task.MockTaskRepository) {
				task := domain_task.ReconstructTask(
					dummyTaskID,
					dummyTaskName,
					dummyUserID,
					domain_task.TaskStatusUnDone,
					0,
					dummyDueDate,
				)
				mockTaskFactory.EXPECT().Create(dummyTaskName, dummyDueDate, dummyUserID).Return(task)
				mockTaskRepository.EXPECT().Insert(gomock.Any(), task).Return(nil)
			},
			args: args{
				input: CreateTaskUseCaseInput{
					TaskName: dummyTaskNameString,
					DueDate:  dummyDueDate,
					UserID:   dummyUserID,
				},
			},
			want: CreateTaskUseCaseDTO{
				TaskID: dummyTaskIDString,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockTaskFactory := mock_task.NewMockTaskFactory(mockCtrl)
			mockTaskRepository := mock_task.NewMockTaskRepository(mockCtrl)
			// 参考記事「gomockを完全に理解する」
			// https://zenn.dev/sanpo_shiho/articles/01da627ead98f5
			tt.prepareMockFn(mockTaskFactory, mockTaskRepository)

			uc := &createTaskUseCase{
				transaction:    transaction.NewNoopTransaction(),
				taskFactory:    mockTaskFactory,
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
