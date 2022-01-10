package user

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/domain/user/mock_user"
	"ddd-sample/infra/in_memory/transaction"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_registerUserUseCase_Execute(t *testing.T) {
	var (
		dummyUserIDString = "12345678901234567890123456"
		dummyUserID, _    = domain.ParseUserID(dummyUserIDString)
		dummyUserID2, _   = domain.ParseUserID("12345678901234567890123457")
		dummyUserName     = "ダミーユーザー"
	)
	type args struct {
		input RegisterUserParam
	}
	tests := []struct {
		name          string
		prepareMockFn func(*mock_user.MockUserFactory, *mock_user.MockUserRepository)
		args          args
		want          RegisterUserUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockUserFactory *mock_user.MockUserFactory, mockUserRepository *mock_user.MockUserRepository) {
				userName, _ := domain.NewUserName(dummyUserName)
				user := domain.ReconstructUser(dummyUserID, userName)
				mockUserFactory.EXPECT().Create(userName).Return(user)
				mockUserRepository.EXPECT().FindByName(gomock.Any(), userName).Return(nil, nil)
				mockUserRepository.EXPECT().Insert(gomock.Any(), domain.ReconstructUser(dummyUserID, userName)).Return(nil)
			},
			args: args{
				input: RegisterUserParam{
					UserName: dummyUserName,
				},
			},
			want:    RegisterUserUseCaseDTO{UserID: dummyUserIDString},
			wantErr: false,
		},
		{
			name: "異常系：ユーザー重複",
			prepareMockFn: func(mockUserFactory *mock_user.MockUserFactory, mockUserRepository *mock_user.MockUserRepository) {
				userName, _ := domain.NewUserName(dummyUserName)
				user := domain.ReconstructUser(dummyUserID, userName)
				mockUserFactory.EXPECT().Create(userName).Return(user)
				duplicateUser := domain.ReconstructUser(dummyUserID2, userName)
				mockUserRepository.EXPECT().FindByName(gomock.Any(), userName).Return(&duplicateUser, nil)
			},
			args: args{
				input: RegisterUserParam{
					UserName: dummyUserName,
				},
			},
			want:    RegisterUserUseCaseDTO{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockUserFactory := mock_user.NewMockUserFactory(mockCtrl)
			mockUserRepository := mock_user.NewMockUserRepository(mockCtrl)
			// 参考記事「gomockを完全に理解する」
			// https://zenn.dev/sanpo_shiho/articles/01da627ead98f5
			tt.prepareMockFn(mockUserFactory, mockUserRepository)

			uc := &registerUserUseCase{
				transaction:            transaction.NewNoopTransaction(),
				userFactory:            mockUserFactory,
				userDuplicationChecker: domain.NewUserDuplicationChecker(mockUserRepository),
				userRepository:         mockUserRepository,
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
