package user

import (
	"errors"
	"net/http"

	"github.com/joesantosio/example-go-project/domain/factory"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
)

func Create(username string, userRepo entity.RepositoryUser, factoryRepo entity.RepositoryFactory) (string, error) {
	// check if the user was created already
	if dbUser, _ := userRepo.GetByUsername(username); dbUser.Username == username {
		return "", httperr.NewError(http.StatusConflict, errors.New("user already exists"))
	}

	id, err := userRepo.Create(username)
	if err != nil {
		return "", err
	}

	err = factory.CreateUserFactories(username, userRepo, factoryRepo)
	if err != nil {
		return "", err
	}

	return id, err
}

func GetByUserID(
	userID string,
	userRepo entity.RepositoryUser,
) (entity.User, error) {
	if len(userID) == 0 {
		return entity.User{}, httperr.NewError(http.StatusBadRequest, errors.New("invalid user"))
	}

	users, err := userRepo.GetByIDs([]string{userID})
	if err != nil {
		return entity.User{}, err
	}

	if len(users) == 0 {
		return entity.User{}, httperr.NewError(http.StatusBadRequest, errors.New("invalid user"))
	}

	return users[0], nil
}
