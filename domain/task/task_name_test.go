package task

import (
	"reflect"
	"testing"
)

func TestNewTaskName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    TaskName
		wantErr bool
	}{
		{
			name: "10文字以下の値を渡すと、正常にインスタンスが生成される",
			args: args{
				name: "1234567890",
			},
			want: TaskName{
				name: "1234567890",
			},
			wantErr: false,
		},
		{
			name: "11文字以下の値を渡すと、エラーが発生する",
			args: args{
				name: "12345678901",
			},
			want:    TaskName{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTaskName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTaskName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
