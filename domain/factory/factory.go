package factory

import (
	"errors"
	"log"
	"net/http"

	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/httperr"
)

var (
	ENABLED_FACTORIES = []string{"iron", "copper", "gold"}
)

func convertModelToFactory(model entity.Factory) *factoryEntity {
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
	userRepo entity.RepositoryUser,
	factoryRepo entity.RepositoryFactory,
) (bool, error) {
	user, err := userRepo.GetUserByUsername(username)
	if err != nil {
		return false, err
	}

	if user == nil {
		return false, httperr.NewError(http.StatusNotFound, errors.New("user not found"))
	}

	factoriesRepo, err := factoryRepo.GetByUsername(username)
	if err != nil {
		return false, err
	}

	// lets first map the resources totals and parse them into
	// an interface that knows how to handle extra stuff
	factories := []*factoryEntity{}
	totals := make(map[string]int)
	for _, res := range factoriesRepo {
		factory := convertModelToFactory(res)
		if factory == nil {
			continue
		}

		totals[res.GetKind()] = res.GetTotal()
		factories = append(factories, factory)
	}

	// lets find the right resource to upgrade
	for _, res := range factories {
		if res.Kind != kind {
			continue
		}

		if res.NextUpgradeCost[kind] > totals[kind] {
			break
		}

		// setup a callback function to save on the repo when the factory
		// is upgraded
		factories[len(factories)-1].updCb = func(total int, level int) {
			_, err := factoryRepo.PatchFactory(kind, username, total, level)

			if err != nil {
				// TODO: should probably have other way to notify the error
				log.Fatalf("ERR: error patching user %s: %v \n", username, err)
			}
		}

		// all in stock, we can update
		res.Upgrade()
	}

	return true, nil
}

func CreateUserFactories(
	username string,
	userRepo entity.RepositoryUser,
	factoryRepo entity.RepositoryFactory,
) error {
	for _, kind := range ENABLED_FACTORIES {
		_, err := factoryRepo.CreateFactory(kind, 0, 0, username)
		if err != nil {
			return err
		}
	}

	return nil
}
