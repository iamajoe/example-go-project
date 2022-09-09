package inmem

import (
	"fmt"
	"math/rand"

	"github.com/joesantosio/example-go-project/entity"
)

// ------------------------------
// repository

type repositoryUser struct {
	data []*entity.User
}

func (repo *repositoryUser) GetByIDs(ids []string) ([]entity.User, error) {
	if len(ids) == 0 {
		return []entity.User{}, nil
	}

	users := []entity.User{}
	for _, id := range ids {
		for _, val := range repo.data {
			if val.ID == id {
				users = append(users, *val)
				break
			}
		}
	}

	return users, nil
}

func (repo *repositoryUser) GetByUsername(username string) (entity.User, error) {
	user := entity.NewModelUser("", "")

	for _, val := range repo.data {
		if val.Username == username {
			user.ID = val.ID
			user.Username = val.Username
			break
		}
	}

	return user, nil
}

func (repo *repositoryUser) Create(username string) (string, error) {
	id := fmt.Sprintf("%d", rand.Intn(10000000000))
	user := entity.NewModelUser(id, username)

	repo.data = append(repo.data, &user)

	return id, nil
}

func createRepositoryUser() (entity.RepositoryUser, error) {
	repo := repositoryUser{
		data: []*entity.User{},
	}

	return &repo, nil
}
