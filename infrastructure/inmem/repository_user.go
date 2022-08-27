package inmem

import "github.com/joesantosio/simple-game-api/infrastructure"

type repositoryUserSingle struct {
	username string
}

type repositoryUser struct {
	data []*repositoryUserSingle
}

func (repo *repositoryUser) GetUserByUsername(username string) (infrastructure.User, error) {
	user := infrastructure.NewUser("")

	for _, val := range repo.data {
		if val.username == username {
			user.SetUsername(val.username)
			break
		}
	}

	return user, nil
}

func (repo *repositoryUser) CreateUser(user infrastructure.User) (bool, error) {
	repo.data = append(repo.data, &repositoryUserSingle{user.GetUsername()})
	return true, nil
}

func createRepositoryUser() (infrastructure.RepositoryUser, error) {
	repo := repositoryUser{
		data: []*repositoryUserSingle{},
	}

	return &repo, nil
}
