package sqlite

import (
	"errors"

	"github.com/joesantosio/example-go-project/entity"
)

type repositories struct {
	db        *DB
	user      entity.RepositoryUser
	userToken entity.RepositoryUserToken
	factory   entity.RepositoryFactory
}

func (r *repositories) GetUser() entity.RepositoryUser {
	return r.user
}

func (r *repositories) GetUserToken() entity.RepositoryUserToken {
	return r.userToken
}

func (r *repositories) GetFactory() entity.RepositoryFactory {
	return r.factory
}

func (r *repositories) Close() error {
	if r.db != nil {
		err := r.db.Close()
		if err != nil {
			return err
		}

		r.db = nil
	}

	return nil
}

func InitRepos(db *DB) (repos entity.Repositories, err error) {
	if db == nil {
		return repos, errors.New("database didn't came in")
	}

	user, err := createRepositoryUser(db)
	if err != nil {
		return repos, err
	}

	userToken, err := createRepositoryUserToken(db)
	if err != nil {
		return repos, err
	}

	factory, err := createRepositoryFactory(db)
	if err != nil {
		return repos, err
	}

	return &repositories{db, user, userToken, factory}, nil
}
