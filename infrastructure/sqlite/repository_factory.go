package sqlite

import (
	"fmt"
	"sync"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

var (
	factoryMutex sync.Mutex
)

type repositoryFactory struct {
	db *DB
}

func (repo *repositoryFactory) GetByUsername(username string) ([]infrastructure.Factory, error) {
	factories := []infrastructure.Factory{}

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

		factory := infrastructure.NewFactory(kind, total, level)
		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) CreateFactory(factory infrastructure.Factory, username string) (bool, error) {
	factoryMutex.Lock()

	sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
	_, err := repo.db.db.Exec(sts, username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
	if err != nil {
		return false, err
	}

	factoryMutex.Unlock()

	return true, err
}

func (repo *repositoryFactory) PatchFactory(factory infrastructure.Factory, username string) (bool, error) {
	factoryMutex.Lock()

	sts := "UPDATE factories SET total=?, level=? WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, factory.GetTotal(), factory.GetLevel(), username, factory.GetKind())
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

func (repo *repositoryFactory) RemoveFactory(factory infrastructure.Factory, username string) (bool, error) {
	sts := "DELETE FROM factories WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, username, factory.GetKind())
	return true, err
}

func createRepositoryFactory(db *DB) (infrastructure.RepositoryFactory, error) {
	repo := repositoryFactory{db}

	factoryMutex.Lock()

	// TODO: should have an id
	// TODO: should have a reference to the users table
	sts := "CREATE TABLE IF NOT EXISTS factories(username TEXT, kind TEXT, total INTEGER, level INTEGER);"
	_, err := repo.db.db.Exec(sts)

	factoryMutex.Unlock()

	return &repo, err
}
