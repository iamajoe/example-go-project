package sqlite

import (
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

	err := repo.db.db.QueryRow("SELECT username FROM users WHERE username=$1", username).Scan(&dbUsername)
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

	// TODO: should have an id
	sts := "CREATE TABLE IF NOT EXISTS users(username TEXT);"
	_, err := repo.db.db.Exec(sts)

	userMutex.Unlock()

	return &repo, err
}
