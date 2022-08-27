package user

import (
	"errors"

	"github.com/joesantosio/simple-game-api/domains/factory"
	"github.com/joesantosio/simple-game-api/infrastructure"
)

func CreateUser(username string, userRepo infrastructure.RepositoryUser, factoryRepo infrastructure.RepositoryFactory) (bool, error) {
	// check if the user was created already
	if dbUser, _ := userRepo.GetUserByUsername(username); dbUser.GetUsername() == username {
		return false, errors.New("user already exists")
	}

	_, err := userRepo.CreateUser(infrastructure.NewUser(username))
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
	userRepo infrastructure.RepositoryUser,
) (infrastructure.User, error) {
	if len(username) == 0 {
		return nil, errors.New("invalid user")
	}

	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	if len(user.GetUsername()) == 0 {
		return user, errors.New("invalid user")
	}

	return user, nil
}
