package sqlite

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/joesantosio/example-go-project/entity"
)

// ------------------------------
// model

// ------------------------------
// repository

type repositoryUser struct {
	tableName string
	db        *DB
}

func (repo *repositoryUser) GetByIDs(ids []string) ([]entity.User, error) {
	if len(ids) == 0 {
		return []entity.User{}, nil
	}

	users := []entity.User{}

	// construct the data for the query
	rawIds := []interface{}{}
	idSql := ""
	for i, id := range ids {
		if len(idSql) != 0 {
			idSql = idSql + " OR"
		}

		idSql = fmt.Sprintf("%s id=%d", idSql, i+1)
		rawIds = append(rawIds, id)
	}

	rows, err := repo.db.db.Query("SELECT id,username FROM "+repo.tableName+" WHERE"+idSql, rawIds...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var username string

		err = rows.Scan(&id, &username)
		if err != nil {
			return nil, err
		}

		user := entity.NewModelUser(id, username)
		users = append(users, user)
	}

	return users, nil
}

func (repo *repositoryUser) GetByUsername(username string) (entity.User, error) {
	user := entity.NewModelUser("", "")

	var dbId string
	var dbUsername string

	err := repo.db.db.QueryRow(
		"SELECT id,username FROM "+repo.tableName+" WHERE username=$1", username,
	).Scan(&dbId, &dbUsername)
	if err != nil {
		return user, err
	}

	user.ID = dbId
	user.Username = dbUsername

	return user, nil
}

func (repo *repositoryUser) Create(username string) (string, error) {
	generatedId := uuid.New().String()

	sts, err := repo.db.db.Prepare("INSERT INTO " + repo.tableName + "(id, username) VALUES(?, ?) RETURNING id")
	if err != nil {
		return "", err
	}
	defer sts.Close()

	var id string
	err = sts.QueryRow(generatedId, username).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, err
}

func (repo *repositoryUser) RemoveTable() (bool, error) {
	sts := "DROP TABLE IF EXISTS " + repo.tableName + " CASCADE"
	_, err := repo.db.db.Exec(sts)
	return true, err
}

func createRepositoryUser(db *DB) (entity.RepositoryUser, error) {
	repo := repositoryUser{"users", db}

	// TODO: should setup migrations
	// TODO: should use tableName
	sts := `
		CREATE TABLE IF NOT EXISTS users(
			id              TEXT			 PRIMARY KEY,
			username 			  TEXT
		);
	`
	_, err := repo.db.db.Exec(sts)

	return &repo, err
}
