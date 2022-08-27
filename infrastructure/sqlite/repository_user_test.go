package sqlite

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

func TestRepositoryUser_GetUserByUsername(t *testing.T) {
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
			sts := "INSERT INTO users VALUES(?);"
			_, err = db.db.Exec(sts, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			// run
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
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

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
			repo, err := createRepositoryUser(db)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.CreateUser(tt.args.user)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got != true {
				t.Errorf("RepositoryUser.CreateUser() = %v, want %v", got, true)
			}

			// check if user is in
			username := ""
			err = db.db.QueryRow(
				fmt.Sprintf("SELECT username FROM users WHERE username='%s'", tt.args.user.GetUsername()),
			).Scan(&username)
			if err != nil {
				t.Fatal(err)
			}

			if username != tt.args.user.GetUsername() {
				t.Errorf("username = %v, want %v", username, tt.args.user.GetUsername())
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
