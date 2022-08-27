package inmem

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

func TestFactoryRepository_GetByUsername(t *testing.T) {
	type fields struct {
		factory infrastructure.Factory
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
				factory: infrastructure.NewFactory(
					fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.CreateFactory(tt.fields.factory, tt.args.username)
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

func TestFactoryRepository_CreateFactory(t *testing.T) {
	type args struct {
		factory  infrastructure.Factory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := infrastructure.NewFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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

func TestFactoryRepository_PatchFactory(t *testing.T) {
	type args struct {
		factory  infrastructure.Factory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := infrastructure.NewFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.CreateFactory(tt.args.factory, tt.args.username)
			if err != nil {
				t.Fatal(err)
			}

			// change total factory by creating a new one
			factory := infrastructure.NewFactory(
				tt.args.factory.GetKind(),
				tt.args.factory.GetLevel(),
				10,
			)

			got, err := repo.PatchFactory(factory, tt.args.username)
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

			if dataGot[0].GetTotal() != factory.GetTotal() {
				t.Errorf("factory.GetTotal() = %v, want %v", dataGot[0].GetTotal(), factory.GetTotal())
			}
		})
	}
}

func TestFactoryRepository_RemoveFactoriesFromUser(t *testing.T) {
	type fields struct {
		factory infrastructure.Factory
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
			factory := infrastructure.NewFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.CreateFactory(tt.fields.factory, tt.args.username)
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
		factory  infrastructure.Factory
		username string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := infrastructure.NewFactory(
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
