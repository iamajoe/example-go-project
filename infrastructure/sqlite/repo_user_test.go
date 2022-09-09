package sqlite

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/joesantosio/example-go-project/entity"
)

func TestRepositoryUser_GetByUsername(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			username := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name: "runs",
				args: args{username},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryUser(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.Create(tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			// run
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
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

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
			repo, err := createRepositoryUser(db)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.Create(tt.args.user.Username)
			if err != nil {
				t.Fatal(err)
				return
			}

			if len(got) == 0 {
				t.Errorf("RepositoryUser.Create() = %v, want %v", got, true)
			}

			// check if user is in
			username := ""
			err = db.db.QueryRow(
				fmt.Sprintf("SELECT username FROM users WHERE username='%s'", tt.args.user.Username),
			).Scan(&username)
			if err != nil {
				t.Fatal(err)
			}

			if username != tt.args.user.Username {
				t.Errorf("username = %v, want %v", username, tt.args.user.Username)
			}
		})
	}
}

func Test_createRepositoryUser(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryUser(db)
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
