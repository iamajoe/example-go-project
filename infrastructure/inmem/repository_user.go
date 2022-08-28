package inmem

import (
	"github.com/joesantosio/simple-game-api/entity"
)

// ------------------------------
// model

type modelUser struct {
	Username string `json:"username"`
}

func (u *modelUser) GetUsername() string {
	return u.Username
}

func newModelUser(username string) *modelUser {
	return &modelUser{username}
}

// ------------------------------
// repository

type repositoryUser struct {
	data []*modelUser
}

func (repo *repositoryUser) GetUserByUsername(username string) (entity.User, error) {
	user := newModelUser("")

	for _, val := range repo.data {
		if val.GetUsername() == username {
			user.Username = val.GetUsername()
			break
		}
	}

	return user, nil
}

func (repo *repositoryUser) CreateUser(username string) (bool, error) {
	repo.data = append(repo.data, &modelUser{username})
	return true, nil
}

func createRepositoryUser() (entity.RepositoryUser, error) {
	repo := repositoryUser{
		data: []*modelUser{},
	}

	return &repo, nil
}
