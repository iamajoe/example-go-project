package factory

func getGoldUpgradeMap() map[int]factoryUpgradeMap {
	upgradeMap := make(map[int]factoryUpgradeMap)

	// level 1
	nextUpgradeCost := make(map[string]int)
	nextUpgradeCost["iron"] = 0
	nextUpgradeCost["copper"] = 100
	nextUpgradeCost["gold"] = 2
	upgradeMap[1] = factoryUpgradeMap{2, 15, nextUpgradeCost}

	// level 2
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 0
	nextUpgradeCost["copper"] = 200
	nextUpgradeCost["gold"] = 4
	upgradeMap[2] = factoryUpgradeMap{3, 30, nextUpgradeCost}

	// level 3
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 0
	nextUpgradeCost["copper"] = 400
	nextUpgradeCost["gold"] = 8
	upgradeMap[3] = factoryUpgradeMap{4, 60, nextUpgradeCost}

	// level 4
	nextUpgradeCost = make(map[string]int)
	nextUpgradeCost["iron"] = 0
	nextUpgradeCost["copper"] = 800
	nextUpgradeCost["gold"] = 16
	upgradeMap[4] = factoryUpgradeMap{6, 90, nextUpgradeCost}

	// level 5
	nextUpgradeCost = make(map[string]int)
	upgradeMap[5] = factoryUpgradeMap{8, 120, make(map[string]int)}

	return upgradeMap
}
