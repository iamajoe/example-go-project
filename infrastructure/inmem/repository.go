package inmem

import (
	"github.com/joesantosio/example-go-project/entity"
)

type repositories struct {
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
	return nil
}

func InitRepos() (repos entity.Repositories, err error) {
	factory, err := createRepositoryFactory()
	if err != nil {
		return repos, err
	}

	userToken, err := createRepositoryUserToken()
	if err != nil {
		return repos, err
	}

	user, err := createRepositoryUser()
	if err != nil {
		return repos, err
	}

	return &repositories{user, userToken, factory}, nil
}
