package entity

type Factory interface {
	GetKind() string
	GetTotal() int
	GetLevel() int
}

type RepositoryFactory interface {
	GetByUsername(username string) ([]Factory, error)
	CreateFactory(kind string, total int, level int, username string) (bool, error)
	PatchFactory(kind string, username string, total int, level int) (bool, error)
	RemoveFactoriesFromUser(username string) (bool, error)
	RemoveFactory(kind string, username string) (bool, error)
}