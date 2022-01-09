package user

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/domain/user/mock_user"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_fetchUserDetailUseCase_Execute(t *testing.T) {
	var (
		dummyUserIDString   = "12345678901234567890123456"
		dummyUserID, _      = domain.ParseUserID(dummyUserIDString)
		dummyUserNameString = "ダミーユーザー"
		dummyUserName, _    = domain.NewUserName(dummyUserNameString)
		dummyUser           = domain.ReconstructUser(dummyUserID, dummyUserName)
	)
	type args struct {
		input FetchUserDetailUseCaseInput
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock_user.MockUserRepository)
		args          args
		want          FetchUserDetailUseCaseDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(gomock.Any(), dummyUserID).Return(&dummyUser, nil)
			},
			args: args{
				input: FetchUserDetailUseCaseInput{
					UserID: dummyUserIDString,
				},
			},
			want: FetchUserDetailUseCaseDTO{
				UserID:   dummyUserIDString,
				UserName: dummyUserNameString,
			},
			wantErr: false,
		},
		{
			name: "異常系：FindByIDエラー",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(gomock.Any(), dummyUserID).Return(nil, errors.New("dummy"))
			},
			args: args{
				input: FetchUserDetailUseCaseInput{
					UserID: dummyUserIDString,
				},
			},
			wantErr: true,
		},
		{
			name: "異常系：ユーザーが存在しない",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().FindByID(gomock.Any(), dummyUserID).Return(nil, nil)
			},
			args: args{
				input: FetchUserDetailUseCaseInput{
					UserID: dummyUserIDString,
				},
			},
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

			uc := &fetchUserDetailUseCase{
				userRepository: mockUserRepository,
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
