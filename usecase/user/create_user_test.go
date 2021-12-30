package user

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/domain/user/mock_user"
	"ddd-sample/usecase/transaction"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_createUserUseCase_Execute(t *testing.T) {
	var (
		dummyUserID   = "" // TODO
		dummyUserName = "ダミーユーザー"
	)
	type args struct {
		input CreateUserUseCaseInput
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock_user.MockUserRepository)
		args          args
		want          CreateUserUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				userName, _ := domain.NewUserName(dummyUserName)
				mockUserRepository.EXPECT().FindByName(userName).Return(nil, nil)
				mockUserRepository.EXPECT().Insert(domain.ReconstructUser(dummyUserID, userName)).Return(nil)
			},
			args: args{
				input: CreateUserUseCaseInput{
					Name: dummyUserName,
				},
			},
			want:    CreateUserUseCaseDTO{},
			wantErr: false,
		},
		{
			name: "異常系：ユーザー重複",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				userName, _ := domain.NewUserName(dummyUserName)
				duplicateUser := domain.ReconstructUser(dummyUserID, userName)
				mockUserRepository.EXPECT().FindByName(userName).Return(&duplicateUser, nil)
			},
			args: args{
				input: CreateUserUseCaseInput{
					Name: dummyUserName,
				},
			},
			want:    CreateUserUseCaseDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			// 参考記事「gomockを完全に理解する」
			// https://zenn.dev/sanpo_shiho/articles/01da627ead98f5
			mockUserRepository := mock_user.NewMockUserRepository(mockCtrl)
			tt.prepareMockFn(mockUserRepository)

			uc := &createUserUseCase{
				userRepository: mockUserRepository,
				transaction:    &transaction.NoopTransaction{},
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
