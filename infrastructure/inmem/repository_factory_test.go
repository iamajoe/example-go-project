package inmem

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFactoryRepository_GetByUsername(t *testing.T) {
	type fields struct {
		factory *modelFactory
	}
	type args struct {
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "runs",
			fields: fields{
				factory: newModelFactory(
					fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
					fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
					10,
					11,
				),
			},
			args: args{fmt.Sprintf("tmp_user_%d", rand.Intn(100000))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.CreateFactory(
				tt.fields.factory.GetKind(), 
				tt.fields.factory.GetTotal(),
				tt.fields.factory.GetLevel(),
				tt.args.username,
			)
			if err != nil {
				t.Fatal(err)
			}

			// run
			got, err := repo.GetByUsername(tt.args.username)
			if len(got) != 1 {
				t.Errorf("FactoryRepository.GetByUsername() = %v, want %v", len(got), 1)
			}

			if got[0].GetKind() != tt.fields.factory.GetKind() {
				t.Errorf("got[0].GetKind() = %v, want %v", got[0].GetKind(), tt.fields.factory.GetKind())
			}
		})
	}
}

func TestFactoryRepository_CreateFactory(t *testing.T) {
	type args struct {
		factory  *modelFactory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := newModelFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory:  factory,
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.CreateFactory(
				tt.args.factory.GetKind(), 
				tt.args.factory.GetTotal(),
				tt.args.factory.GetLevel(),
				tt.args.username,
			)
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

			if dataGot[0].GetKind() != tt.args.factory.GetKind() {
				t.Errorf("got[0].GetKind() = %v, want %v", dataGot[0].GetKind(), tt.args.factory.GetKind())
			}
		})
	}
}

func TestFactoryRepository_PatchFactory(t *testing.T) {
	type args struct {
		factory  *modelFactory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := newModelFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory:  factory,
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.CreateFactory(
				tt.args.factory.GetKind(), 
				tt.args.factory.GetTotal(), 
				tt.args.factory.GetLevel(), 
				tt.args.username,
			)
			if err != nil {
				t.Fatal(err)
			}

			
			newTotal := 10

			got, err := repo.PatchFactory(
				tt.args.factory.GetKind(), 
				tt.args.username,
				newTotal, 
				tt.args.factory.GetLevel(), 
			)
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

			if dataGot[0].GetTotal() != newTotal {
				t.Errorf("newTotal = %v, want %v", dataGot[0].GetTotal(), newTotal)
			}
		})
	}
}

func TestFactoryRepository_RemoveFactoriesFromUser(t *testing.T) {
	type fields struct {
		factory *modelFactory
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
			factory := newModelFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				fields: fields{
					factory: factory,
				},
				args: args{
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.CreateFactory(
				tt.fields.factory.GetKind(), 
				tt.fields.factory.GetTotal(), 
				tt.fields.factory.GetLevel(), 
				tt.args.username,
			)
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

func TestFactoryRepository_RemoveFactory(t *testing.T) {
	type args struct {
		factory  *modelFactory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := newModelFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory:  factory,
					username: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				},
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			_, err = repo.CreateFactory(
				tt.args.factory.GetKind(), 
				tt.args.factory.GetTotal(), 
				tt.args.factory.GetLevel(), 
				tt.args.username,
			)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.RemoveFactory(tt.args.factory.GetKind(), tt.args.username)
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

func Test_createRepositoryFactory(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"runs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryFactory()
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryFactory() = %v, want %v", got, nil)
			}
		})
	}
}
