package entity

type Repositories interface {
	Close() error
	GetUser() RepositoryUser
	GetFactory() RepositoryFactory
	GetUserToken() RepositoryUserToken
}
