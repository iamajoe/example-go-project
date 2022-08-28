package entity

type User interface {
	GetUsername() string
}

type RepositoryUser interface {
	GetUserByUsername(username string) (User, error)
	CreateUser(username string) (bool, error)
}