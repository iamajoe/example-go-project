package main

import (
	"fmt"
	"sync"
)

var (
	factorySqliteMutex sync.Mutex
)

type repositoryFactorySqlite struct {
	db *DBSqlite
}

func (repo *repositoryFactorySqlite) GetByUsername(username string) ([]Factory, error) {
	factories := []Factory{}

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

		var factory Factory
		switch kind {
		case "iron":
			factory = &IronFactory{
				Kind:  kind,
				Total: total,
				Level: level,
			}
		case "copper":
			factory = &CopperFactory{
				Kind:  kind,
				Total: total,
				Level: level,
			}
		case "gold":
			factory = &GoldFactory{
				Kind:  kind,
				Total: total,
				Level: level,
			}
		}

		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactorySqlite) CreateFactory(factory Factory, username string) (bool, error) {
	factorySqliteMutex.Lock()

	sts := "INSERT INTO factories VALUES(?, ?, ?, ?);"
	_, err := repo.db.db.Exec(sts, username, factory.GetKind(), factory.GetTotal(), factory.GetLevel())
	if err != nil {
		return false, err
	}

	factorySqliteMutex.Unlock()

	return true, err
}

func (repo *repositoryFactorySqlite) PatchFactory(factory Factory, username string) (bool, error) {
	factorySqliteMutex.Lock()

	sts := "UPDATE factories SET total=?, level=? WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, factory.GetTotal(), factory.GetLevel(), username, factory.GetKind())
	if err != nil {
		return false, err
	}

	factorySqliteMutex.Unlock()

	return true, err
}

func (repo *repositoryFactorySqlite) RemoveFactoriesFromUser(username string) (bool, error) {
	sts := "DELETE FROM factories WHERE username=?"
	_, err := repo.db.db.Exec(sts, username)
	return true, err
}

func (repo *repositoryFactorySqlite) RemoveFactory(factory Factory, username string) (bool, error) {
	sts := "DELETE FROM factories WHERE username=? AND kind=?"
	_, err := repo.db.db.Exec(sts, username, factory.GetKind())
	return true, err
}

func createRepositoryFactorySqlite(db *DBSqlite) (RepositoryFactory, error) {
	repo := repositoryFactorySqlite{db}

	factorySqliteMutex.Lock()

	// TODO: should have an id
	// TODO: should have a reference to the users table
	sts := "CREATE TABLE IF NOT EXISTS factories(username TEXT, kind TEXT, total INTEGER, level INTEGER);"
	_, err := repo.db.db.Exec(sts)

	factorySqliteMutex.Unlock()

	return &repo, err
}
