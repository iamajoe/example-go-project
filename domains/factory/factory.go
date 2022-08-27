package factory

import (
	"errors"
	"log"

	"github.com/joesantosio/simple-game-api/infrastructure"
)

func convertModelToFactory(model infrastructure.Factory) *Factory {
	var upgradeMap map[int]factoryUpgradeMap
	switch model.GetKind() {
	case "copper":
		upgradeMap = getCopperUpgradeMap()
	case "gold":
		upgradeMap = getGoldUpgradeMap()
	case "iron":
		upgradeMap = getIronUpgradeMap()
	}

	if upgradeMap == nil {
		return nil
	}

	return newFactory(model.GetKind(), model.GetTotal(), model.GetLevel(), upgradeMap)
}

func UpgradeUserResource(
	username string,
	kind string,
	userRepo infrastructure.RepositoryUser,
	factoryRepo infrastructure.RepositoryFactory,
) (bool, error) {
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, errors.New("user not found")
	}

	factoriesRepo, err := factoryRepo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	// lets first map the resources totals and parse them into
	// an interface that knows how to handle extra stuff
	factories := []*Factory{}
	totals := make(map[string]int)
	for _, res := range factoriesRepo {
		factory := convertModelToFactory(res)
		if factory == nil {
			continue
		}

		totals[res.GetKind()] = res.GetTotal()
		factories = append(factories, factory)
	}

	// setup a callback function to save on the repo when the factory
	// is upgraded
	updCb := func() {
		factories, err := factoryRepo.GetByUsername(username)
		if err != nil {
			// TODO: should probably have other way to notify the error
			log.Fatalf("ERR: error fetching factories %s: %v \n", username, err)
			return
		}

		// save the user with the new data on the resources
		for _, factory := range factories {
			_, err := factoryRepo.PatchFactory(factory, username)

			if err != nil {
				// TODO: should probably have other way to notify the error
				log.Fatalf("ERR: error patching user %s: %v \n", username, err)
			}
		}
	}

	// lets find the right resource to upgrade
	for _, res := range factories {
		if res.Kind != kind {
			continue
		}

		if res.NextUpgradeCost[kind] > totals[kind] {
			break
		}

		factories[len(factories)-1].updCb = updCb

		// all in stock, we can update
		res.Upgrade()
	}

	return true, nil
}

func CreateUserFactories(
	username string,
	userRepo infrastructure.RepositoryUser,
	factoryRepo infrastructure.RepositoryFactory,
) error {
	for _, kind := range []string{"iron", "copper", "gold"} {
		factoryModel := infrastructure.NewFactory(kind, 0, 0)
		_, err := factoryRepo.CreateFactory(factoryModel, username)
		if err != nil {
			return err
		}
	}

	return nil
}
