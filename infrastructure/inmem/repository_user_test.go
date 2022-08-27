package inmem

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

func TestRepositoryUser_GetUserByUsername(t *testing.T) {
	type fields struct {
		data []*repositoryUserSingle
	}
	type args struct {
		username string
	}
	type testStruct struct {
		name   string
		fields fields
		args   args
	}

	tests := []testStruct{
		func() testStruct {
			username := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name: "runs",
				fields: fields{
					data: []*repositoryUserSingle{{username}},
				},
				args: args{username},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repositoryUser{
				data: tt.fields.data,
			}
			got, err := repo.GetUserByUsername(tt.args.username)

			if err != nil {
				t.Fatal(err)
				return
			}

			if got.GetUsername() != tt.args.username {
				t.Errorf("got.GetUsername() = %v, want %v", got.GetUsername(), tt.args.username)
			}
		})
	}
}

func TestRepositoryUser_CreateUser(t *testing.T) {
	type args struct {
		user infrastructure.User
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			username := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))
			user := infrastructure.NewUser(username)

			return testStruct{
				name: "runs",
				args: args{user},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repositoryUser{}

			got, err := repo.CreateUser(tt.args.user)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got != true {
				t.Errorf("RepositoryUser.CreateUser() = %v, want %v", got, true)
			}

			// run
			gotUser, err := repo.GetUserByUsername(tt.args.user.GetUsername())
			if gotUser.GetUsername() != tt.args.user.GetUsername() {
				t.Errorf("RepositoryUser.GetByUsername() = %v, want %v", gotUser.GetUsername(), tt.args.user.GetUsername())
			}
		})
	}
}

func Test_createRepositoryUser(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryUser()
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryUser() = %v, want %v", got, nil)
			}
		})
	}
}
