package sqlite

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func Test_repositoryUserToken_Create(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		userId string
		token  string
	}
	type testStruct struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}

	tests := []testStruct{
		func() testStruct {
			userId := fmt.Sprintf("%d", rand.Intn(100000))
			token := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name:    "runs",
				args:    args{userId, token},
				want:    true,
				wantErr: false,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := createRepositoryUserToken(db)
			if err != nil {
				t.Fatal(err)
			}

			got, err := r.Create(tt.args.userId, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("repositoryUserToken.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repositoryUserToken.Create() = %v, want %v", got, tt.want)
			}

			if got {
				var dbExists bool
				sql := "SELECT EXISTS (SELECT id FROM usertokens WHERE token=$1 AND userid=$2)"
				err := db.db.QueryRow(sql, tt.args.token, tt.args.userId).Scan(&dbExists)
				if err != nil {
					t.Fatal(err)
				}

				if !dbExists {
					t.Errorf("found = %v, want %v", dbExists, true)
				}
			}
		})
	}
}

func Test_repositoryUserToken_IsTokenValid(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type args struct {
		userID string
		token  string
	}
	type testStruct struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}

	tests := []testStruct{
		func() testStruct {
			token := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name:    "runs invalid",
				args:    args{"", token},
				want:    false,
				wantErr: false,
			}
		}(),

		func() testStruct {
			userID := fmt.Sprintf("%d", rand.Intn(100000))
			token := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name:    "runs valid",
				args:    args{userID, token},
				want:    true,
				wantErr: false,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := createRepositoryUserToken(db)
			if err != nil {
				t.Fatal(err)
			}

			// prepare by creating the user on the db
			if len(tt.args.userID) > 0 {
				fmt.Println("creating!!", tt.args.userID)
				_, err = r.Create(tt.args.userID, tt.args.token)
				if err != nil {
					t.Fatal(err)
				}
			}

			got, err := r.IsTokenValid(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("repositoryUserToken.IsTokenValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repositoryUserToken.IsTokenValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createRepositoryUserToken(t *testing.T) {
	path := fmt.Sprintf("tmp_test_%d.db", rand.Intn(10000))
	db, err := Connect(path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	defer os.Remove(path)

	type testStruct struct {
		name string
	}

	tests := []testStruct{
		func() testStruct {
			return testStruct{
				name: "runs",
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createRepositoryUserToken(db)
			if err != nil {
				t.Errorf("createRepositoryUserToken() error = %v, wantErr %v", err, nil)
				return
			}
			if got == nil {
				t.Errorf("createRepositoryUserToken() = %v, want %v", got, "repo")
			}
		})
	}
}
