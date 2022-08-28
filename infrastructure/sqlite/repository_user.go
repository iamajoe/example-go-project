package sqlite

import (
	"fmt"
	"sync"

	"github.com/joesantosio/simple-game-api/entity"
)

var (
	userMutex sync.Mutex
)

// ------------------------------
// model

type modelUser struct {
	Username string `json:"username"`
}

func (u *modelUser) GetUsername() string {
	return u.Username
}

func newModelUser(username string) *modelUser {
	return &modelUser{username}
}

// ------------------------------
// repository

type repositoryUser struct {
	db *DB
}

func (repo *repositoryUser) GetUserByUsername(username string) (entity.User, error) {
	user := newModelUser("")

	var dbUsername string

	// TODO: shouldnt use sprintf but a lib to make sure that we don't have security issues
	err := repo.db.db.QueryRow(
		fmt.Sprintf("SELECT username FROM users WHERE username='%s'", username),
	).Scan(&dbUsername)
	if err != nil {
		return user, err
	}

	user.Username = dbUsername

	return user, nil
}

func (repo *repositoryUser) CreateUser(username string) (bool, error) {
	userMutex.Lock()

	sts := "INSERT INTO users VALUES(?);"
	_, err := repo.db.db.Exec(sts, username)
	if err != nil {
		return false, err
	}

	userMutex.Unlock()

	return true, err
}

func createRepositoryUser(db *DB) (entity.RepositoryUser, error) {
	repo := repositoryUser{db}

	userMutex.Lock()

	// TODO: should have an id but for the assessment, no need
	sts := "CREATE TABLE IF NOT EXISTS users(username TEXT);"
	_, err := repo.db.db.Exec(sts)

	userMutex.Unlock()

	return &repo, err
}
