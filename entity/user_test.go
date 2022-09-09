package entity

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_newModelUserToken(t *testing.T) {
	type args struct {
		id     int
		userId string
		token  string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			return testStruct{
				name: "runs",
				args: args{
					id:     rand.Intn(100000),
					userId: fmt.Sprintf("%d", rand.Intn(100000)),
					token:  fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewModelUserToken(tt.args.id, tt.args.userId, tt.args.token)

			if got.UserId != tt.args.userId {
				t.Errorf("got.userId = %v, want %v", got.UserId, tt.args.userId)
			}

			if got.Token != tt.args.token {
				t.Errorf("got.token = %v, want %v", got.Token, tt.args.token)
			}
		})
	}
}
