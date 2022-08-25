package main

type repositoryUserInmemSingle struct {
	username string
}

type RepositoryUserInmem struct {
	data              []*repositoryUserInmemSingle
	factoryRepository RepositoryFactory
}

func (repo *RepositoryUserInmem) GetUserByUsername(username string) (User, error) {
	user := User{"", []Factory{}, repo, repo.factoryRepository}

	for _, val := range repo.data {
		if val.username == username {
			user.Username = val.username
			break
		}
	}

	return user, nil
}

func (repo *RepositoryUserInmem) CreateUser(user User) (bool, error) {
	repo.data = append(repo.data, &repositoryUserInmemSingle{user.GetUsername()})
	return true, nil
}

func createRepositoryUserInmem(factoryRepo RepositoryFactory) (RepositoryUser, error) {
	repositoryUser := RepositoryUserInmem{
		data:              []*repositoryUserInmemSingle{},
		factoryRepository: factoryRepo,
	}

	return &repositoryUser, nil
}
