package user

import (
	"errors"
	"net/http"

	"github.com/joesantosio/simple-game-api/domain/factory"
	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/httperr"
)

func CreateUser(username string, userRepo entity.RepositoryUser, factoryRepo entity.RepositoryFactory) (bool, error) {
	// check if the user was created already
	if dbUser, _ := userRepo.GetUserByUsername(username); dbUser.GetUsername() == username {
		return false, httperr.NewError(http.StatusConflict, errors.New("user already exists"))
	}

	_, err := userRepo.CreateUser(username)
	if err != nil {
		return false, err
	}

	err = factory.CreateUserFactories(username, userRepo, factoryRepo)
	if err != nil {
		return false, err
	}

	return true, err
}

func GetUserByUsername(
	username string,
	userRepo entity.RepositoryUser,
) (entity.User, error) {
	if len(username) == 0 {
		return nil, httperr.NewError(http.StatusBadRequest, errors.New("invalid user"))
	}

	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	if len(user.GetUsername()) == 0 {
		return user, httperr.NewError(http.StatusBadRequest, errors.New("invalid user"))
	}

	return user, nil
}
