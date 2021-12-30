package user

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/domain/user/mock_user"
	"ddd-sample/usecase/transaction"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_updateUserUseCase_Execute(t *testing.T) {
	var (
		dummyUserID      = "1"
		dummyUserName, _ = domain.NewUserName("ダミーユーザー")
		dummyUser        = domain.ReconstructUser(dummyUserID, dummyUserName)
	)
	type args struct {
		input UpdateUserUseCaseInput
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock_user.MockUserRepository)
		args          args
		want          UpdateUserUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(dummyUserID).Return(&dummyUser, nil)
				userName, _ := domain.NewUserName("ダミーユーザー2")
				mockUserRepository.EXPECT().FindByName(userName).Return(nil, nil)
				mockUserRepository.EXPECT().Update(domain.ReconstructUser(dummyUserID, userName)).Return(nil)
			},
			args: args{
				input: UpdateUserUseCaseInput{
					ID:   dummyUserID,
					Name: "ダミーユーザー2",
				},
			},
			want:    UpdateUserUseCaseDTO{},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが存在しない",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(dummyUserID).Return(nil, nil)
			},
			args: args{
				input: UpdateUserUseCaseInput{
					ID:   dummyUserID,
					Name: "ダミーユーザー2",
				},
			},
			want:    UpdateUserUseCaseDTO{},
			wantErr: true,
		},
		{
			name: "異常系：ユーザー重複",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(dummyUserID).Return(&dummyUser, nil)
				userName, _ := domain.NewUserName("ダミーユーザー2")
				duplicateUser := domain.ReconstructUser(dummyUserID, userName)
				mockUserRepository.EXPECT().FindByName(userName).Return(&duplicateUser, nil)
			},
			args: args{
				input: UpdateUserUseCaseInput{
					ID:   dummyUserID,
					Name: "ダミーユーザー2",
				},
			},
			want:    UpdateUserUseCaseDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUserRepository := mock_user.NewMockUserRepository(mockCtrl)
			// 参考記事「gomockを完全に理解する」
			// https://zenn.dev/sanpo_shiho/articles/01da627ead98f5
			tt.prepareMockFn(mockUserRepository)

			uc := &updateUserUseCase{
				transaction:    &transaction.NoopTransaction{},
				userRepository: mockUserRepository,
			}
			got, err := uc.Execute(tt.args.input)
			if (err != nil) != tt.wantErr { // エラーがある、かつ期待結果と一致しない場合
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
