package factory

func getCopperUpgradeMap() map[int]factoryUpgradeMap {
	upgradeMap := make(map[int]factoryUpgradeMap)

	// level 1
	nextUpgradeCost := make(map[string]int)
	nextUpgradeCost["iron"] = 300
	nextUpgradeCost["copper"] = 100
	nextUpgradeCost["gold"] = 1
	upgradeMap[1] = factoryUpgradeMap{10, 15, nextUpgradeCost}

	// level 2
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 800
	nextUpgradeCost["copper"] = 250
	nextUpgradeCost["gold"] = 2
	upgradeMap[2] = factoryUpgradeMap{20, 30, nextUpgradeCost}

	// level 3
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 1600
	nextUpgradeCost["copper"] = 500
	nextUpgradeCost["gold"] = 4
	upgradeMap[3] = factoryUpgradeMap{40, 60, nextUpgradeCost}

	// level 4
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 3000
	nextUpgradeCost["copper"] = 1000
	nextUpgradeCost["gold"] = 8
	upgradeMap[4] = factoryUpgradeMap{80, 90, nextUpgradeCost}

	// level 5
	nextUpgradeCost = make(map[string]int)
	upgradeMap[5] = factoryUpgradeMap{150, 120, make(map[string]int)}

	return upgradeMap
}
