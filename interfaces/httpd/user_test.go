package httpd

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joesantosio/simple-game-api/domains/user"
	"github.com/joesantosio/simple-game-api/infrastructure/inmem"
)

func TestReqCreateUser(t *testing.T) {
	repos, err := inmem.InitRepos()
	if err != nil {
		t.Fatal(err)
	}
	defer repos.Close()

	username := fmt.Sprintf("%d", rand.Intn(10000))
	body := []byte(fmt.Sprintf(`{"username":"%s"}`, username))

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(reqCreateUser(repos))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("wrong status code: got %v want %v", rec.Code, http.StatusOK)
	}

	if rec.Body.String() != "{\"ok\":true,\"code\":200,\"data\":\"true\"}" {
		t.Errorf("body = %v, want %v", rec.Body.String(), "true")
	}

	user, err := repos.User.GetUserByUsername(username)
	if err != nil {
		t.Fatal(err)
	}

	if len(user.GetUsername()) == 0 {
		t.Errorf("length = %v, want %v", len(user.GetUsername()), 0)
	}

	if user.GetUsername() != username {
		t.Errorf("username = %v, want %v", user.GetUsername(), username)
	}
}

func TestReqGetDashboard(t *testing.T) {
	repos, err := inmem.InitRepos()
	if err != nil {
		t.Fatal(err)
	}
	defer repos.Close()

	username := fmt.Sprintf("%d", rand.Intn(10000))
	_, err = user.CreateUser(username, repos.User, repos.Factory)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/dashboard?username=%s", username), nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(reqGetDashboard(repos))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("wrong status code: got %v want %v", rec.Code, http.StatusOK)
	}

	expected := fmt.Sprintf("{\"ok\":true,\"code\":200,\"data\":{\"username\":\"%s\",\"factories\":[{\"isRunning\":false,\"kind\":\"iron\",\"total\":0,\"level\":1,\"isUpgrade\":false,\"productionTime\":0,\"nextUpgradeTime\":0,\"nextUpgradeCost\":null},{\"isRunning\":false,\"kind\":\"copper\",\"total\":0,\"level\":1,\"isUpgrade\":false,\"productionTime\":0,\"nextUpgradeTime\":0,\"nextUpgradeCost\":null},{\"isRunning\":false,\"kind\":\"gold\",\"total\":0,\"level\":1,\"isUpgrade\":false,\"productionTime\":0,\"nextUpgradeTime\":0,\"nextUpgradeCost\":null}]}}", username)
	if rec.Body.String() != expected {
		t.Errorf("body = %v, want %v", rec.Body.String(), expected)
	}
}

// TODO: should test that the main is actually creating a server
