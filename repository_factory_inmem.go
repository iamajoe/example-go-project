package main

type repositoryFactoryInmemSingle struct {
	username string
	kind     string
	total    int
	level    int
}

type repositoryFactoryInmem struct {
	data []repositoryFactoryInmemSingle
}

func (repo *repositoryFactoryInmem) GetByUsername(username string) ([]Factory, error) {
	factories := []Factory{}
	for _, val := range repo.data {
		if val.username != username {
			continue
		}

		var factory Factory
		switch val.kind {
		case "iron":
			factory = &IronFactory{
				Kind:  val.kind,
				Total: val.total,
				Level: val.level,
			}
		case "copper":
			factory = &CopperFactory{
				Kind:  val.kind,
				Total: val.total,
				Level: val.level,
			}
		case "gold":
			factory = &GoldFactory{
				Kind:  val.kind,
				Total: val.total,
				Level: val.level,
			}
		}

		factories = append(factories, factory)
	}

	return factories, nil
}

func (repo *repositoryFactoryInmem) CreateFactory(factory Factory, username string) (bool, error) {
	repo.data = append(repo.data, repositoryFactoryInmemSingle{
		username: username,
		kind:     factory.GetKind(),
		total:    factory.GetTotal(),
		level:    factory.GetLevel(),
	})

	return true, nil
}

func (repo *repositoryFactoryInmem) PatchFactory(factory Factory, username string) (bool, error) {
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

func (repo *repositoryFactoryInmem) RemoveFactoriesFromUser(username string) (bool, error) {
	newData := []repositoryFactoryInmemSingle{}

	for _, val := range repo.data {
		if val.username == username {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func (repo *repositoryFactoryInmem) RemoveFactory(factory Factory, username string) (bool, error) {
	newData := []repositoryFactoryInmemSingle{}

	for _, val := range repo.data {
		if val.username == username && val.kind == factory.GetKind() {
			continue
		}

		newData = append(newData, val)
	}

	repo.data = newData

	return true, nil
}

func createRepositoryFactoryInmem() (RepositoryFactory, error) {
	repo := repositoryFactoryInmem{
		data: []repositoryFactoryInmemSingle{},
	}

	return &repo, nil
}
