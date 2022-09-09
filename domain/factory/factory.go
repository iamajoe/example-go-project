package factory

import (
	"errors"
	"log"
	"net/http"

	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
)

var (
	ENABLED_FACTORIES = []string{"iron", "copper", "gold"}
)

func convertModelToFactory(model entity.Factory) *factoryEntity {
	var upgradeMap map[int]factoryUpgradeMap
	switch model.Kind {
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

	return newFactory(model.Kind, model.Total, model.Level, upgradeMap)
}

func UpgradeUserResource(
	userID string,
	kind string,
	userRepo entity.RepositoryUser,
	factoryRepo entity.RepositoryFactory,
) (bool, error) {
	users, err := userRepo.GetByIDs([]string{userID})
	if err != nil {
		return false, err
	}

	if len(users) == 0 {
		return false, httperr.NewError(http.StatusNotFound, errors.New("user not found"))
	}

	factoriesRepo, err := factoryRepo.GetByUserID(userID)
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

		totals[res.Kind] = res.Total
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
			_, err := factoryRepo.Patch(kind, userID, total, level)

			if err != nil {
				// TODO: should probably have other way to notify the error
				log.Fatalf("ERR: error patching user %s: %v \n", userID, err)
			}
		}

		// all in stock, we can update
		res.Upgrade()
	}

	return true, nil
}

func CreateUserFactories(
	userID string,
	userRepo entity.RepositoryUser,
	factoryRepo entity.RepositoryFactory,
) error {
	for _, kind := range ENABLED_FACTORIES {
		_, err := factoryRepo.Create(kind, 0, 0, userID)
		if err != nil {
			return err
		}
	}

	return nil
}
