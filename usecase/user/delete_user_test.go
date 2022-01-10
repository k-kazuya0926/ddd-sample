package user

import (
	domain "ddd-sample/domain/user"
	"ddd-sample/domain/user/mock_user"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_deleteUserUseCase_Execute(t *testing.T) {
	var (
		dummyUserIDString = "12345678901234567890123456"
		dummyUserID, _    = domain.ParseUserID(dummyUserIDString)
	)
	type args struct {
		input DeleteUserParam
	}
	tests := []struct {
		name          string
		prepareMockFn func(m *mock_user.MockUserRepository)
		args          args
		want          DeleteUserDTO
		wantErr       bool
	}{
		{
			name: "正常系",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().Delete(gomock.Any(), dummyUserID).Return(nil)
			},
			args: args{
				input: DeleteUserParam{
					UserID: dummyUserIDString,
				},
			},
			want:    DeleteUserDTO{},
			wantErr: false,
		},
		{
			name: "異常系：ユーザーが存在しない",
			prepareMockFn: func(mockUserRepository *mock_user.MockUserRepository) {
				mockUserRepository.EXPECT().Delete(gomock.Any(), dummyUserID).Return(errors.New("dummy"))
			},
			args: args{
				input: DeleteUserParam{
					UserID: dummyUserIDString,
				},
			},
			want:    DeleteUserDTO{},
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

			uc := &deleteUserUseCase{
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
