package inmem

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joesantosio/example-go-project/entity"
)

func TestRepositoryUser_GetByUsername(t *testing.T) {
	type fields struct {
		data []*entity.User
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
			id := fmt.Sprintf("tmp_id_%d", rand.Intn(100000))
			username := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name: "runs",
				fields: fields{
					data: []*entity.User{{id, username}},
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
			got, err := repo.GetByUsername(tt.args.username)

			if err != nil {
				t.Fatal(err)
				return
			}

			if got.Username != tt.args.username {
				t.Errorf("got.Username = %v, want %v", got.Username, tt.args.username)
			}
		})
	}
}

func TestRepositoryUser_Create(t *testing.T) {
	type args struct {
		user entity.User
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			username := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))
			user := entity.NewModelUser("", username)

			return testStruct{
				name: "runs",
				args: args{user},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repositoryUser{}

			got, err := repo.Create(tt.args.user.Username)
			if err != nil {
				t.Fatal(err)
				return
			}

			if len(got) == 0 {
				t.Errorf("RepositoryUser.CreateUser() = %v, want %v", got, true)
			}

			// run
			gotUser, err := repo.GetByUsername(tt.args.user.Username)
			if gotUser.Username != tt.args.user.Username {
				t.Errorf("RepositoryUser.GetByUsername() = %v, want %v", gotUser.Username, tt.args.user.Username)
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
