package inmem

import "github.com/joesantosio/simple-game-api/entity"

// ------------------------------
// model

type modelFactory struct {
	username string
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

func newModelFactory(username string, kind string, total int, level int) *modelFactory {
	return &modelFactory{
		username: username,
		Kind:  kind,
		Total: total,
		Level: level,
	}
}

// ------------------------------
// repository

type repositoryFactory struct {
	data []*modelFactory
}

func (repo *repositoryFactory) GetByUsername(username string) ([]entity.Factory, error) {
	factories := []entity.Factory{}
	for _, val := range repo.data {
		if val.username != username {
			continue
		}

		factory := newModelFactory(username, val.Kind, val.Total, val.Level)
		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) CreateFactory(kind string, total int, level int, username string) (bool, error) {
	repo.data = append(repo.data, newModelFactory(username, kind, total, level))
	return true, nil
}

func (repo *repositoryFactory) PatchFactory(kind string, username string, total int, level int) (bool, error) {
	for _, val := range repo.data {
		if val.username != username && val.Kind != kind {
			continue
		}

		val.Kind = kind
		val.Total = total
		val.Level = level
	}

	return true, nil
}

func (repo *repositoryFactory) RemoveFactoriesFromUser(username string) (bool, error) {
	newData := []*modelFactory{}

	for _, val := range repo.data {
		if val.username == username {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func (repo *repositoryFactory) RemoveFactory(kind string, username string) (bool, error) {
	newData := []*modelFactory{}

	for _, val := range repo.data {
		if val.username == username && val.Kind == kind {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func createRepositoryFactory() (entity.RepositoryFactory, error) {
	repo := repositoryFactory{
		data: []*modelFactory{},
	}

	return &repo, nil
}
