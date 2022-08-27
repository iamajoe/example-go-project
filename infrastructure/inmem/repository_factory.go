package inmem

import "github.com/joesantosio/simple-game-api/infrastructure"

type repositoryFactorySingle struct {
	username string
	kind     string
	total    int
	level    int
}

type repositoryFactory struct {
	data []*repositoryFactorySingle
}

func (repo *repositoryFactory) GetByUsername(username string) ([]infrastructure.Factory, error) {
	factories := []infrastructure.Factory{}
	for _, val := range repo.data {
		if val.username != username {
			continue
		}

		factory := infrastructure.NewFactory(val.kind, val.total, val.level)
		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) CreateFactory(factory infrastructure.Factory, username string) (bool, error) {
	repo.data = append(repo.data, &repositoryFactorySingle{
		username: username,
		kind:     factory.GetKind(),
		total:    factory.GetTotal(),
		level:    factory.GetLevel(),
	})

	return true, nil
}

func (repo *repositoryFactory) PatchFactory(factory infrastructure.Factory, username string) (bool, error) {
	for _, val := range repo.data {
		if val.username != username {
			continue
		}

		val.kind = factory.GetKind()
		val.total = factory.GetTotal()
		val.level = factory.GetLevel()
	}

	return true, nil
}

func (repo *repositoryFactory) RemoveFactoriesFromUser(username string) (bool, error) {
	newData := []*repositoryFactorySingle{}

	for _, val := range repo.data {
		if val.username == username {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func (repo *repositoryFactory) RemoveFactory(factory infrastructure.Factory, username string) (bool, error) {
	newData := []*repositoryFactorySingle{}

	for _, val := range repo.data {
		if val.username == username && val.kind == factory.GetKind() {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func createRepositoryFactory() (infrastructure.RepositoryFactory, error) {
	repo := repositoryFactory{
		data: []*repositoryFactorySingle{},
	}

	return &repo, nil
}
