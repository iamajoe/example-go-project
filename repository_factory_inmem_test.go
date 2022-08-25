package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFactoryRepositoryInmem_GetByUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		{"runs", args{fmt.Sprintf("tmp_user_%d", rand.Intn(100000))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			_, err = repo.CreateFactory(&factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			// run
			got, err := repo.GetByUsername(tt.args.username)
			if len(got) != 1 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(got), 1)
			}

			if got[0].GetKind() != "copper" {
				t.Errorf("got[0].GetKind() = %v, want %v", got[0].GetKind(), "copper")
			}
		})
	}
}

func TestFactoryRepositoryInmem_CreateFactory(t *testing.T) {
	type args struct {
		factory  Factory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := newCopperFactory(func() {})

			return testStruct{
				name: "runs",
				args: args{
					factory:  &factory,
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.CreateFactory(tt.args.factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			if got != true {
				t.Errorf("FactoryRepository.CreateFactory() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUsername(tt.args.username)
			if len(dataGot) != 1 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(dataGot), 1)
			}
			if err != nil {
				t.Fatal(err)
			}

			if dataGot[0].GetKind() != "copper" {
				t.Errorf("got[0].GetKind() = %v, want %v", dataGot[0].GetKind(), "copper")
			}
		})
	}
}

func TestFactoryRepositoryInmem_PatchFactory(t *testing.T) {
	type args struct {
		factory  Factory
		username string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			_, err = repo.CreateFactory(&factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			// change total factory
			factory.Total = 10

			got, err := repo.PatchFactory(tt.args.factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			if got != true {
				t.Errorf("FactoryRepository.PatchFactory() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUsername(tt.args.username)
			if len(dataGot) != 1 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(dataGot), 1)
			}
			if err != nil {
				t.Fatal(err)
			}

			if dataGot[0].GetTotal() != factory.Total {
				t.Errorf("factory.Total = %v, want %v", dataGot[0].GetTotal(), factory.Total)
			}
		})
	}
}

func TestFactoryRepositoryInmem_RemoveFactoriesFromUser(t *testing.T) {
	type args struct {
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "runs",
			args: args{
				username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			_, err = repo.CreateFactory(&factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.RemoveFactoriesFromUser(tt.args.username)
			if err != nil {
				t.Fatal(err)
			}
			if got != true {
				t.Errorf("FactoryRepository.RemoveFactoriesFromUser() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUsername(tt.args.username)
			if len(dataGot) != 0 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(dataGot), 0)
			}
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFactoryRepositoryInmem_RemoveFactory(t *testing.T) {
	type args struct {
		factory  Factory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := newCopperFactory(func() {})

			return testStruct{
				name: "runs",
				args: args{
					factory:  &factory,
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.CreateFactory(tt.args.factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.RemoveFactory(tt.args.factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}
			if got != true {
				t.Errorf("FactoryRepository.RemoveFactory() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUsername(tt.args.username)
			if len(dataGot) != 0 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(dataGot), 0)
			}
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func Test_createRepositoryFactoryInmem(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryFactoryInmem()
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryFactoryInmem() = %v, want %v", got, nil)
			}
		})
	}
}
