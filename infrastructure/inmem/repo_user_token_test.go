package inmem

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joesantosio/example-go-project/entity"
)

func Test_repositoryUserToken_Create(t *testing.T) {
	type fields struct {
		data []*entity.UserToken
	}
	type args struct {
		userId string
		token  string
	}
	type testStruct struct {
		name    string
		fields  fields
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
				fields:  fields{[]*entity.UserToken{}},
				args:    args{userId, token},
				want:    true,
				wantErr: false,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repositoryUserToken{
				data: tt.fields.data,
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
				found := false
				for _, u := range r.data {
					if u.UserId == tt.args.userId && u.Token == tt.args.token {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("found = %v, want %v", true, found)
				}
			}
		})
	}
}

func Test_repositoryUserToken_IsTokenValid(t *testing.T) {
	type fields struct {
		data []*entity.UserToken
	}
	type args struct {
		token string
	}
	type testStruct struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}

	tests := []testStruct{
		func() testStruct {
			token := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name:    "runs invalid",
				fields:  fields{[]*entity.UserToken{}},
				args:    args{token},
				want:    false,
				wantErr: false,
			}
		}(),

		func() testStruct {
			id := rand.Intn(100000)
			userId := fmt.Sprintf("%d", rand.Intn(100000))
			token := fmt.Sprintf("tmp_user_%d", rand.Intn(100000))

			return testStruct{
				name:    "runs valid",
				fields:  fields{[]*entity.UserToken{{id, userId, token}}},
				args:    args{token},
				want:    true,
				wantErr: false,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repositoryUserToken{
				data: tt.fields.data,
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
			got, err := createRepositoryUserToken()
			if err != nil {
				t.Errorf("createRepositoryUser() error = %v, wantErr %v", err, nil)
				return
			}
			if got == nil {
				t.Errorf("createRepositoryUser() = %v, want %v", got, "repo")
			}
		})
	}
}
