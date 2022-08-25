package main

type RepositoryFactory interface {
	GetByUsername(username string) ([]Factory, error)
	CreateFactory(factory Factory, username string) (bool, error)
	PatchFactory(factory Factory, username string) (bool, error)
	RemoveFactoriesFromUser(username string) (bool, error)
	RemoveFactory(factory Factory, username string) (bool, error)
}
