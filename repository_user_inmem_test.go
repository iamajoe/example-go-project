package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepositoryUserInmem_GetUserByUsername(t *testing.T) {
	type fields struct {
		data []*repositoryUserInmemSingle
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
					data: []*repositoryUserInmemSingle{{username}},
				},
				args: args{username},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factoryRepo, _ := createRepositoryFactoryInmem()
			repo := &RepositoryUserInmem{
				data:              tt.fields.data,
				factoryRepository: factoryRepo,
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

func TestRepositoryUserInmem_CreateUser(t *testing.T) {
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
			factoryRepo, _ := createRepositoryFactoryInmem()

			repo := &RepositoryUserInmem{
				factoryRepository: factoryRepo,
			}
			tt.args.user.factoryRepository = factoryRepo
			tt.args.user.userRepository = repo

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
			if gotUser.Username != tt.args.user.GetUsername() {
				t.Errorf("RepositoryUser.GetByUsername() = %v, want %v", gotUser.Username, tt.args.user.GetUsername())
			}
		})
	}
}

func Test_createRepositoryUserInmem(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factoryRepo, _ := createRepositoryFactoryInmem()
			got, err := createRepositoryUserInmem(factoryRepo)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryUserInmem() = %v, want %v", got, nil)
			}
		})
	}
}
