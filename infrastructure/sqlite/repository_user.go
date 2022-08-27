package sqlite

import (
	"fmt"
	"sync"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

var (
	userMutex sync.Mutex
)

type repositoryUser struct {
	db *DB
}

func (repo *repositoryUser) GetUserByUsername(username string) (infrastructure.User, error) {
	user := infrastructure.NewUser("")

	var dbUsername string

	// TODO: shouldnt use sprintf but a lib to make sure that we don't have security issues
	err := repo.db.db.QueryRow(
		fmt.Sprintf("SELECT username FROM users WHERE username='%s'", username),
	).Scan(&dbUsername)
	if err != nil {
		return user, err
	}

	user.SetUsername(dbUsername)

	return user, nil
}

func (repo *repositoryUser) CreateUser(user infrastructure.User) (bool, error) {
	userMutex.Lock()

	sts := "INSERT INTO users VALUES(?);"
	_, err := repo.db.db.Exec(sts, user.GetUsername())
	if err != nil {
		return false, err
	}

	userMutex.Unlock()

	return true, err
}

func createRepositoryUser(db *DB) (infrastructure.RepositoryUser, error) {
	repo := repositoryUser{db}

	userMutex.Lock()

	// TODO: should have an id but for the assessment, no need
	sts := "CREATE TABLE IF NOT EXISTS users(username TEXT);"
	_, err := repo.db.db.Exec(sts)

	userMutex.Unlock()

	return &repo, err
}
