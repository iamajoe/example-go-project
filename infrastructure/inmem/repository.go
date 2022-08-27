package inmem

import "github.com/joesantosio/simple-game-api/infrastructure"

func InitRepos() (repos *infrastructure.Repositories, err error) {
	closeFn := func() error {
		return nil
	}

	factory, err := createRepositoryFactory()
	if err != nil {
		return repos, err
	}

	user, err := createRepositoryUser()
	if err != nil {
		return repos, err
	}

	repos = infrastructure.NewRepositories(user, factory, closeFn)

	return repos, nil
}
