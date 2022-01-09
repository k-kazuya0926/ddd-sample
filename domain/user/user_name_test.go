package user

import (
	"reflect"
	"testing"
)

func TestNewUserName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr bool
	}{
		{
			name: "正常系：3文字",
			args: args{
				name: "あいう",
			},
			want: UserName{
				name: "あいう",
			},
			wantErr: false,
		},
		{
			name: "異常系：2文字",
			args: args{
				name: "あい",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
