package sqlite

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/joesantosio/simple-game-api/entity"
)

func TestFactoryRepository_GetByUsername(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type fields struct {
		factory entity.Factory
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
					10,
					11,
				),
			},
			args: args{fmt.Sprintf("tmp_user_%d", rand.Intn(100000))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactory(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, tt.fields.factory.GetKind(), tt.fields.factory.GetTotal(), tt.fields.factory.GetLevel())
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
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		factory  entity.Factory
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
			repo, err := createRepositoryFactory(db)
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
			count := 0
			rows, err := db.db.Query(
				fmt.Sprintf("SELECT kind FROM factories WHERE username='%s' AND kind='%s'", tt.args.username, tt.args.factory.GetKind()),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				count += 1
			}

			if count != 1 {
				t.Errorf("count on db = %v, want %v", count, 1)
			}
		})
	}
}

func TestFactoryRepository_PatchFactory(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		factory  entity.Factory
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
			repo, err := createRepositoryFactory(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, tt.args.factory.GetKind(), tt.args.factory.GetTotal(), tt.args.factory.GetLevel())
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
			count := 0
			rows, err := db.db.Query(
				fmt.Sprintf("SELECT total FROM factories WHERE username='%s' AND kind='%s'", tt.args.username, tt.args.factory.GetKind()),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				count += 1

				var total int

				err = rows.Scan(&total)
				if err != nil {
					t.Fatal(err)
				}

				if total != newTotal {
					t.Errorf("factory.Total = %v, want %v", total, newTotal)
				}
			}

			if count != 1 {
				t.Errorf("count on db = %v, want %v", count, 1)
			}
		})
	}
}

func TestFactoryRepository_RemoveFactoriesFromUser(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type fields struct {
		factory entity.Factory
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
			repo, err := createRepositoryFactory(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, tt.fields.factory.GetKind(), tt.fields.factory.GetTotal(), tt.fields.factory.GetLevel())
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
			count := 0
			rows, err := db.db.Query(
				fmt.Sprintf("SELECT kind FROM factories WHERE username='%s' AND kind='%s'", tt.args.username, tt.fields.factory.GetKind()),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				count += 1
			}

			if count != 0 {
				t.Errorf("count on db = %v, want %v", count, 0)
			}
		})
	}
}

func TestFactoryRepository_RemoveFactory(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		factory  entity.Factory
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
			repo, err := createRepositoryFactory(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, tt.args.factory.GetKind(), tt.args.factory.GetTotal(), tt.args.factory.GetLevel())
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
			count := 0
			rows, err := db.db.Query(
				fmt.Sprintf("SELECT kind FROM factories WHERE username='%s' AND kind='%s'", tt.args.username, tt.args.factory.GetKind()),
			)
			if err != nil {
				t.Fatal(err)
			}
			defer rows.Close()
			for rows.Next() {
				count += 1
			}

			if count != 0 {
				t.Errorf("count on db = %v, want %v", count, 0)
			}
		})
	}
}

func Test_createRepositoryFactory(t *testing.T) {
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
			got, err := createRepositoryFactory(db)
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
