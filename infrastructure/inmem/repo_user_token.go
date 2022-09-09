package inmem

import (
	"github.com/joesantosio/example-go-project/entity"
)

// ------------------------------
// repository

type repositoryUserToken struct {
	data []*entity.UserToken
}

func (r *repositoryUserToken) Create(userId string, token string) (bool, error) {
	newData := []*entity.UserToken{}
	for _, val := range r.data {
		if val.UserId == userId {
			continue
		}

		newData = append(newData, val)
	}

	userToken := entity.NewModelUserToken(len(r.data), userId, token)

	newData = append(newData, &userToken)
	r.data = newData

	return true, nil
}

func (r *repositoryUserToken) IsTokenValid(token string) (bool, error) {
	for _, val := range r.data {
		if val.Token == token {
			return true, nil
		}
	}

	return false, nil
}

func createRepositoryUserToken() (entity.RepositoryUserToken, error) {
	repo := repositoryUserToken{
		data: []*entity.UserToken{},
	}

	return &repo, nil
}
