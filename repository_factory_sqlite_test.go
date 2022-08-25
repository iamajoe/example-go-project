package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestFactoryRepositorySqlite_GetByUsername(t *testing.T) {
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
	tests := []struct {
		name string
		args args
	}{
		{"runs", args{fmt.Sprintf("tmp_user_%d", rand.Intn(100000))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := createRepositoryFactorySqlite(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
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

func TestFactoryRepositorySqlite_CreateFactory(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
	defer os.Remove(path)

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
			repo, err := createRepositoryFactorySqlite(db)
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

func TestFactoryRepositorySqlite_PatchFactory(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
	defer os.Remove(path)

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
			repo, err := createRepositoryFactorySqlite(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
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

				if total != factory.Total {
					t.Errorf("factory.Total = %v, want %v", total, factory.Total)
				}
			}

			if count != 1 {
				t.Errorf("count on db = %v, want %v", count, 1)
			}
		})
	}
}

func TestFactoryRepositorySqlite_RemoveFactoriesFromUser(t *testing.T) {
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
			repo, err := createRepositoryFactorySqlite(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
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
				fmt.Sprintf("SELECT kind FROM factories WHERE username='%s' AND kind='%s'", tt.args.username, factory.GetKind()),
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

func TestFactoryRepositorySqlite_RemoveFactory(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := ConnectSQLite(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.db.Close()
	defer os.Remove(path)

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
			repo, err := createRepositoryFactorySqlite(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare
			factory := newCopperFactory(func() {})
			sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
			_, err = db.db.Exec(sts, tt.args.username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
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

func Test_createRepositoryFactorySqlite(t *testing.T) {
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
			got, err := createRepositoryFactorySqlite(db)
			if err != nil {
				t.Fatal(err)
				return
			}

			if got == nil {
				t.Errorf("createRepositoryFactorySqlite() = %v, want %v", got, nil)
			}
		})
	}
}
