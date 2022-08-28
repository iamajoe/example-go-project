package sqlite

import (
	"fmt"
	"sync"

	"github.com/joesantosio/simple-game-api/entity"
)

var (
	factoryMutex sync.Mutex
)

// ------------------------------
// model

type modelFactory struct {
	Kind  string `json:"kind"`
	Total int    `json:"total"`
	Level int    `json:"level"`
}

func (r *modelFactory) GetKind() string {
	return r.Kind
}
func (r *modelFactory) GetTotal() int {
	return r.Total
}
func (r *modelFactory) GetLevel() int {
	return r.Level
}

func newModelFactory(kind string, total int, level int) *modelFactory {
	return &modelFactory{
		Kind:  kind,
		Total: total,
		Level: level,
	}
}

// ------------------------------
// repository

type repositoryFactory struct {
	db *DB
}

func (repo *repositoryFactory) GetByUsername(username string) ([]entity.Factory, error) {
	factories := []entity.Factory{}

	// TODO: shouldnt use sprintf but a lib to make sure that we don't have security issues
	rows, err := repo.db.db.Query(
		fmt.Sprintf("SELECT kind,total,level FROM factories WHERE username='%s'", username),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var kind string
		var total int
		var level int

		err = rows.Scan(&kind, &total, &level)
		if err != nil {
			return nil, err
		}

		factory := newModelFactory(kind, total, level)
		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) CreateFactory(kind string, total int, level int, username string) (bool, error) {
	factoryMutex.Lock()

	sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
	_, err := repo.db.db.Exec(sts, username, kind, total, level)
	if err != nil {
		return false, err
	}

	factoryMutex.Unlock()

	return true, err
}

func (repo *repositoryFactory) PatchFactory(kind string, username string, total int, level int) (bool, error) {
	factoryMutex.Lock()

	sts := "UPDATE factories SET total=?, level=? WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, total, level, username, kind)
	if err != nil {
		return false, err
	}

	factoryMutex.Unlock()

	return true, err
}

func (repo *repositoryFactory) RemoveFactoriesFromUser(username string) (bool, error) {
	sts := "DELETE FROM factories WHERE username=?"
	_, err := repo.db.db.Exec(sts, username)
	return true, err
}

func (repo *repositoryFactory) RemoveFactory(kind string, username string) (bool, error) {
	sts := "DELETE FROM factories WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, username, kind)
	return true, err
}

func createRepositoryFactory(db *DB) (entity.RepositoryFactory, error) {
	repo := repositoryFactory{db}

	factoryMutex.Lock()

	// TODO: should have an id
	// TODO: should have a reference to the users table
	sts := "CREATE TABLE IF NOT EXISTS factories(username TEXT, kind TEXT, total INTEGER, level INTEGER);"
	_, err := repo.db.db.Exec(sts)

	factoryMutex.Unlock()

	return &repo, err
}
