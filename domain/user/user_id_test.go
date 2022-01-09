package user

import (
	"reflect"
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestParseUserID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    UserID
		wantErr bool
	}{
		{
			name: "正常系：26文字",
			args: args{
				id: "1234567890123456789012345A",
			},
			want: UserID{
				id: ulid.MustParse("1234567890123456789012345A"),
			},
			wantErr: false,
		},
		{
			name: "異常系：25文字",
			args: args{
				id: "1234567890123456789012345",
			},
			wantErr: true,
		},
		{
			name: "異常系：27文字",
			args: args{
				id: "123456789012345678901234567",
			},
			wantErr: true,
		},
		{
			name: "異常系：空文字",
			args: args{
				id: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
