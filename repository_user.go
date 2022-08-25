package main

type RepositoryUser interface {
	GetUserByUsername(username string) (User, error)
	CreateUser(user User) (bool, error)
}
