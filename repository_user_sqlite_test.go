package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestRepositoryUserSqlite_GetUserByUsername(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
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
			repo, err := createRepositoryUserSqlite(db)
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

func TestRepositoryUserSqlite_CreateUser(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
	defer os.Remove(path)

	type args struct {
		user User
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			user := User{
				Username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
			}

			return testStruct{
				name: "runs",
				args: args{user},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryUserSqlite(db)
			if err != nil {
				t.Fatal(err)
			}
			tt.args.user.userRepository = repo

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

			if username != tt.args.user.Username {
				t.Errorf("username = %v, want %v", username, tt.args.user.Username)
			}
		})
	}
}

func Test_createRepositoryUserSqlite(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
	defer os.Remove(path)

	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryUserSqlite(db)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryUserSqlite() = %v, want %v", got, nil)
			}
		})
	}
}
