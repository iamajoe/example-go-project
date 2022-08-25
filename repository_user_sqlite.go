package main

import (
	"fmt"
	"sync"
)

var (
	userSqliteMutex sync.Mutex
)

type RepositoryUserSqlite struct {
	db DBSqlite
}

func (repo *RepositoryUserSqlite) GetUserByUsername(username string) (User, error) {
	var user User

	// TODO: shouldnt use sprintf but a lib to make sure that we don't have security issues
	// TODO: also... not really needed?!
	err := repo.db.db.QueryRow(
		fmt.Sprintf("SELECT username FROM users WHERE username='%s'", username),
	).Scan(&user.Username)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (repo *RepositoryUserSqlite) CreateUser(user User) (bool, error) {
	userSqliteMutex.Lock()

	sts := "INSERT INTO users VALUES(?);"
	_, err := repo.db.db.Exec(sts, user.GetUsername())
	if err != nil {
		return false, err
	}

	userSqliteMutex.Unlock()

	return true, err
}

func createRepositoryUserSqlite(db DBSqlite) (RepositoryUser, error) {
	repositoryUser := RepositoryUserSqlite{db}

	userSqliteMutex.Lock()

	// TODO: should have an id but for the assessment, no need
	sts := "CREATE TABLE IF NOT EXISTS users(username TEXT);"
	_, err := repositoryUser.db.db.Exec(sts)

	userSqliteMutex.Unlock()

	return &repositoryUser, err
}
