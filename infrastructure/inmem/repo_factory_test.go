package inmem

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joesantosio/example-go-project/entity"
)

func TestFactoryRepository_GetByUserID(t *testing.T) {
	type fields struct {
		factory *entity.Factory
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "runs",
			fields: fields{
				factory: entity.NewModelFactory(
					rand.Intn(100000),
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
			_, err = repo.Create(
				tt.fields.factory.Kind,
				tt.fields.factory.Total,
				tt.fields.factory.Level,
				tt.args.userID,
			)
			if err != nil {
				t.Fatal(err)
			}

			// run
			got, err := repo.GetByUserID(tt.args.userID)
			if len(got) != 1 {
				t.Errorf("FactoryRepository.GetByUserID() = %v, want %v", len(got), 1)
			}

			if got[0].Kind != tt.fields.factory.Kind {
				t.Errorf("got[0].Kind = %v, want %v", got[0].Kind, tt.fields.factory.Kind)
			}
		})
	}
}

func TestFactoryRepository_Create(t *testing.T) {
	type args struct {
		factory *entity.Factory
		userID  string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := entity.NewModelFactory(
				rand.Intn(100000),
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory: factory,
					userID:  fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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

			got, err := repo.Create(
				tt.args.factory.Kind,
				tt.args.factory.Total,
				tt.args.factory.Level,
				tt.args.userID,
			)
			if err != nil {
				t.Fatal(err)
			}

			if got != true {
				t.Errorf("FactoryRepository.Create() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUserID(tt.args.userID)
			if len(dataGot) != 1 {
				t.Errorf("FactoryRepository.GetByUserID() = %v, want %v", len(dataGot), 1)
			}
			if err != nil {
				t.Fatal(err)
			}

			if dataGot[0].Kind != tt.args.factory.Kind {
				t.Errorf("got[0].Kind = %v, want %v", dataGot[0].Kind, tt.args.factory.Kind)
			}
		})
	}
}

func TestFactoryRepository_Patch(t *testing.T) {
	type args struct {
		factory *entity.Factory
		userID  string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := entity.NewModelFactory(
				rand.Intn(100000),
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory: factory,
					userID:  fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.Create(
				tt.args.factory.Kind,
				tt.args.factory.Total,
				tt.args.factory.Level,
				tt.args.userID,
			)
			if err != nil {
				t.Fatal(err)
			}

			newTotal := 10

			got, err := repo.Patch(
				tt.args.factory.Kind,
				tt.args.userID,
				newTotal,
				tt.args.factory.Level,
			)
			if err != nil {
				t.Fatal(err)
			}

			if got != true {
				t.Errorf("FactoryRepository.Patch() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUserID(tt.args.userID)
			if len(dataGot) != 1 {
				t.Errorf("FactoryRepository.GetByUserID() = %v, want %v", len(dataGot), 1)
			}
			if err != nil {
				t.Fatal(err)
			}

			if dataGot[0].Total != newTotal {
				t.Errorf("newTotal = %v, want %v", dataGot[0].Total, newTotal)
			}
		})
	}
}

func TestFactoryRepository_RemoveFactoriesFromUser(t *testing.T) {
	type fields struct {
		factory *entity.Factory
	}
	type args struct {
		userID string
	}
	type testStruct struct {
		name   string
		fields fields
		args   args
	}

	tests := []testStruct{
		func() testStruct {
			factory := entity.NewModelFactory(
				rand.Intn(100000),
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
					userID: fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.Create(
				tt.fields.factory.Kind,
				tt.fields.factory.Total,
				tt.fields.factory.Level,
				tt.args.userID,
			)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.RemoveFactoriesFromUser(tt.args.userID)
			if err != nil {
				t.Fatal(err)
			}
			if got != true {
				t.Errorf("FactoryRepository.RemoveFactoriesFromUser() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUserID(tt.args.userID)
			if len(dataGot) != 0 {
				t.Errorf("FactoryRepository.GetByUserID() = %v, want %v", len(dataGot), 0)
			}
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFactoryRepository_Remove(t *testing.T) {
	type args struct {
		factory *entity.Factory
		userID  string
	}
	type testStruct struct {
		name string
		args args
	}

	tests := []testStruct{
		func() testStruct {
			factory := entity.NewModelFactory(
				rand.Intn(100000),
				fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
				fmt.Sprintf("tmp_kind_%d", rand.Intn(100000)),
				10,
				11,
			)

			return testStruct{
				name: "runs",
				args: args{
					factory: factory,
					userID:  fmt.Sprintf("tmp_user_%d", rand.Intn(100000)),
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
			_, err = repo.Create(
				tt.args.factory.Kind,
				tt.args.factory.Total,
				tt.args.factory.Level,
				tt.args.userID,
			)
			if err != nil {
				t.Fatal(err)
			}

			got, err := repo.Remove(tt.args.factory.Kind, tt.args.userID)
			if err != nil {
				t.Fatal(err)
			}
			if got != true {
				t.Errorf("FactoryRepository.Remove() = %v, want %v", got, true)
			}

			// check if factory is in
			dataGot, err := repo.GetByUserID(tt.args.userID)
			if len(dataGot) != 0 {
				t.Errorf("FactoryRepository.GetByUserID() = %v, want %v", len(dataGot), 0)
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
