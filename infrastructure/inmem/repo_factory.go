package inmem

import "github.com/joesantosio/example-go-project/entity"

// ------------------------------
// repository

type repositoryFactory struct {
	data []*entity.Factory
}

func (repo *repositoryFactory) GetByUserID(userID string) ([]entity.Factory, error) {
	factories := []entity.Factory{}
	for _, val := range repo.data {
		if val.UserID != userID {
			continue
		}

		factory := entity.NewModelFactory(val.ID, val.UserID, val.Kind, val.Total, val.Level)
		factories = append(factories, *factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) Create(kind string, total int, level int, userID string) (bool, error) {
	id := len(repo.data)
	factory := entity.NewModelFactory(id, userID, kind, total, level)

	repo.data = append(repo.data, factory)

	return true, nil
}

func (repo *repositoryFactory) Patch(kind string, userID string, total int, level int) (bool, error) {
	for _, val := range repo.data {
		if val.UserID != userID && val.Kind != kind {
			continue
		}

		val.Kind = kind
		val.Total = total
		val.Level = level
	}

	return true, nil
}

func (repo *repositoryFactory) RemoveFactoriesFromUser(userID string) (bool, error) {
	newData := []*entity.Factory{}

	for _, val := range repo.data {
		if val.UserID == userID {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func (repo *repositoryFactory) Remove(kind string, userID string) (bool, error) {
	newData := []*entity.Factory{}

	for _, val := range repo.data {
		if val.UserID == userID && val.Kind == kind {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func createRepositoryFactory() (entity.RepositoryFactory, error) {
	repo := repositoryFactory{
		data: []*entity.Factory{},
	}

	return &repo, nil
}
