package serverstd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joesantosio/example-go-project/acl"
	"github.com/joesantosio/example-go-project/domain/user"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/infrastructure/inmem"
)

func Test_reqCreateUser(t *testing.T) {
	type response struct {
		Ok   bool   `json:"ok"`
		Code int    `json:"code"`
		Data string `json:"data,omitempty"`
		Err  string `json:"err,omitempty"`
	}

	type args struct {
		username string
		body     []byte
		repos    entity.Repositories
	}
	type testStruct struct {
		name     string
		args     args
		wantCode int
		wantBody response
		wantUser bool
	}

	tests := []testStruct{
		func() testStruct {
			repos, err := inmem.InitRepos()
			if err != nil {
				t.Fatal(err)
			}

			username := fmt.Sprintf("%d", rand.Intn(10000))
			_, err = repos.GetUser().Create(username)
			if err != nil {
				t.Fatal(err)
			}

			body := []byte(fmt.Sprintf(`{"username":"%s"}`, username))

			return testStruct{
				name:     "errors if user already exists",
				args:     args{username, body, repos},
				wantCode: http.StatusConflict,
				wantBody: response{
					Ok:   false,
					Code: http.StatusConflict,
					Err:  http.StatusText(http.StatusConflict),
				},
				wantUser: true,
			}
		}(),
		func() testStruct {
			repos, err := inmem.InitRepos()
			if err != nil {
				t.Fatal(err)
			}

			username := fmt.Sprintf("%d", rand.Intn(10000))
			body := []byte(fmt.Sprintf(`{"username":"%s"}`, username))

			return testStruct{
				name:     "runs",
				args:     args{username, body, repos},
				wantCode: http.StatusOK,
				wantBody: response{
					Ok:   true,
					Code: http.StatusOK,
					Err:  "",
					Data: "something...",
				},
				wantUser: true,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.args.repos.Close()

			req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(tt.args.body))
			if err != nil {
				t.Fatal(err)
			}

			rec := httptest.NewRecorder()
			handler := http.HandlerFunc(reqCreateUser(tt.args.repos))
			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("wrong status code: got %v want %v", rec.Code, tt.wantCode)
			}

			var res response
			err = json.NewDecoder(rec.Body).Decode(&res)
			if err != nil {
				t.Fatal(err)
			}

			if res.Ok != tt.wantBody.Ok || res.Err != tt.wantBody.Err || res.Code != tt.wantBody.Code {
				t.Errorf("body = %v, want %v", rec.Body.String(), tt.wantBody)
				return
			}

			if (len(res.Data) == 0 && len(tt.wantBody.Data) > 0) || (len(res.Data) > 0 && len(tt.wantBody.Data) == 0) {
				t.Errorf("data = %v, want %v", res.Data, tt.wantBody.Data)
			}

			user, err := tt.args.repos.GetUser().GetByUsername(tt.args.username)
			if err != nil {
				t.Fatal(err)
			}
			if len(user.Username) == 0 {
				t.Fatalf("len(user.Username): got %v want %v", len(user.Username), 1)
			}
		})
	}
}

func Test_reqGetDashboard(t *testing.T) {
	type response struct {
		Ok   bool        `json:"ok"`
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
		Err  string      `json:"err,omitempty"`
	}

	repos, err := inmem.InitRepos()
	if err != nil {
		t.Fatal(err)
	}
	defer repos.Close()

	username := fmt.Sprintf("%d", rand.Intn(10000))
	id, err := user.Create(username, repos.GetUser(), repos.GetFactory())
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/dashboard", nil)
	ctx, err := acl.SetContextRoles(req.Context(), id, repos)
	if err != nil {
		t.Fatal(err)
	}
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(reqGetDashboard(repos))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("wrong status code: got %v want %v", rec.Code, http.StatusOK)
	}

	var res response
	err = json.NewDecoder(rec.Body).Decode(&res)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Ok || res.Err != "" || res.Code != http.StatusOK {
		t.Errorf("body = %v not right", res)
		return
	}

	if res.Data == nil {
		t.Errorf("data = %v not right", res.Data)
	}
}
