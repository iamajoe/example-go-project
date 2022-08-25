package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username          string    `json:"username"`
	Factories         []Factory `json:"factories"`
	userRepository    RepositoryUser
	factoryRepository RepositoryFactory
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) UpgradeResource(kind string) (bool, error) {
	user, err := getUserByUsername(u.Username, u.userRepository, u.factoryRepository)
	if err != nil {
		return false, err
	}

	// lets first map the resources totals
	totals := make(map[string]int)
	for _, res := range user.Factories {
		totals[res.GetKind()] = res.GetTotal()
	}

	// lets find the right resource to upgrade
	for _, res := range user.Factories {
		if res.GetKind() != kind {
			continue
		}

		if res.GetNextUpgradeCost()[kind] > totals[kind] {
			break
		}

		// all in stock, we can update
		res.Upgrade()
	}

	return true, nil
}

func createUser(username string, userRepository RepositoryUser, factoryRepository RepositoryFactory) (bool, error) {
	// check if the user was created already
	if dbUser, _ := userRepository.GetUserByUsername(username); dbUser.Username == username {
		return false, errors.New("user already exists")
	}

	var user User
	updCb := func() {
		// save the user with the new data on the resources
		for _, factory := range user.Factories {
			_, err := factoryRepository.PatchFactory(factory, user.GetUsername())

			if err != nil {
				// TODO: should probably have other way to notify the error
				fmt.Errorf("error patching user: %v", err)
			}
		}
	}

	IronFactory := newIronFactory(updCb)
	CopperFactory := newCopperFactory(updCb)
	GoldFactory := newGoldFactory(updCb)

	user = User{
		Username:       username,
		Factories:      []Factory{&IronFactory, &CopperFactory, &GoldFactory},
		userRepository: userRepository,
	}

	_, err := userRepository.CreateUser(user)
	if err != nil {
		return false, err
	}

	for _, factory := range user.Factories {
		_, err := factoryRepository.CreateFactory(factory, user.GetUsername())
		if err != nil {
			return false, err
		}
	}

	// don't set the loop when we run the tests, we don't have a way to handle it
	if strings.Contains(os.Args[0], "/_test/") {
		// now that we have the user created, we want to start the resource timers
		for _, res := range user.Factories {
			go res.Start()
		}
	}

	return true, err
}

func getUserByUsername(username string, userRepository RepositoryUser, factoryRepository RepositoryFactory) (User, error) {
	if len(username) == 0 {
		return User{}, errors.New("invalid user")
	}

	user, err := userRepository.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	if len(user.GetUsername()) == 0 {
		return user, errors.New("invalid user")
	}

	user.Factories, err = factoryRepository.GetByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
