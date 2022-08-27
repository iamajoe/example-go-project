package sqlite

import (
	"errors"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

func InitRepos(db *DB) (repos *infrastructure.Repositories, err error) {
	if db == nil {
		return repos, errors.New("database didn't came in")
	}

	closeFn := func() error {
		err := db.Close()
		if err != nil {
			return err
		}

		return nil
	}

	user, err := createRepositoryUser(db)
	if err != nil {
		return repos, err
	}

	factory, err := createRepositoryFactory(db)
	if err != nil {
		return repos, err
	}

	repos = infrastructure.NewRepositories(user, factory, closeFn)

	return repos, nil
}
