package sqlite

import (
	"github.com/joesantosio/example-go-project/entity"
)

// ------------------------------
// repository

type repositoryUserToken struct {
	tableName string
	db        *DB
}

func (repo *repositoryUserToken) Create(userID string, token string) (bool, error) {
	sts := "INSERT INTO " + repo.tableName + "(userid, token) VALUES(?, ?)"
	_, err := repo.db.db.Exec(sts, userID, token)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryUserToken) IsTokenValid(token string) (bool, error) {
	var dbExists bool
	sql := "SELECT EXISTS (SELECT id FROM " + repo.tableName + " WHERE token=$1)"
	err := repo.db.db.QueryRow(sql, token).Scan(&dbExists)
	if err != nil {
		return false, err
	}

	return dbExists, nil
}

func (repo *repositoryUserToken) RemoveTable() (bool, error) {
	sts := "DROP TABLE IF EXISTS " + repo.tableName + " CASCADE"
	_, err := repo.db.db.Exec(sts)
	return true, err
}

func createRepositoryUserToken(db *DB) (entity.RepositoryUserToken, error) {
	repo := repositoryUserToken{"usertokens", db}

	// TODO: missing references
	// TODO: should setup migrations
	// TODO: should use tableName
	sts := `
		CREATE TABLE IF NOT EXISTS usertokens(
			id              INTEGER  PRIMARY KEY AUTOINCREMENT,
			userid 			    TEXT     NOT NULL,
			token 			    TEXT     NOT NULL
		);
	`
	_, err := repo.db.db.Exec(sts)

	return &repo, err
}
