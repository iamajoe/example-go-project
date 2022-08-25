package main

import "time"

type GoldFactory struct {
	IsRunning       bool           `json:"isRunning"`
	Kind            string         `json:"kind"`
	Total           int            `json:"total"`
	Level           int            `json:"level"`
	IsUpgrade       bool           `json:"isUpgrade"`
	Production      int            `json:"productionTime"`
	NextUpgradeTime time.Duration  `json:"nextUpgradeTime"`
	NextUpgradeCost map[string]int `json:"nextUpgradeCost"`
	updCb           func()
}

func (r *GoldFactory) GetKind() string {
	return r.Kind
}

func (r *GoldFactory) GetTotal() int {
	return r.Total
}

func (r *GoldFactory) GetLevel() int {
	return r.Level
}

func (r *GoldFactory) GetNextUpgradeCost() map[string]int {
	return r.NextUpgradeCost
}

func (r *GoldFactory) Start() {
	r.IsRunning = true

	if r.Level == 0 {
		r.UpgradeToLevel(1)
	} else {
		r.UpgradeToLevel(r.Level)
	}

	r.Loop()
}

func (r *GoldFactory) Stop() {
	r.IsRunning = false
}

func (r *GoldFactory) Loop() {
	if !r.IsRunning {
		return
	}

	time.Sleep(time.Second)
	r.Total += r.Production
	r.Loop()
}

func (r *GoldFactory) UpgradeToLevel(level int) {
	switch level {
	case 1:
		r.Production = 2
		r.NextUpgradeTime = 15
		r.NextUpgradeCost["iron"] = 0
		r.NextUpgradeCost["copper"] = 100
		r.NextUpgradeCost["gold"] = 2
	case 2:
		r.Production = 3
		r.NextUpgradeTime = 30
		r.NextUpgradeCost["iron"] = 0
		r.NextUpgradeCost["copper"] = 200
		r.NextUpgradeCost["gold"] = 4
	case 3:
		r.Production = 4
		r.NextUpgradeTime = 60
		r.NextUpgradeCost["iron"] = 0
		r.NextUpgradeCost["copper"] = 400
		r.NextUpgradeCost["gold"] = 8
	case 4:
		r.Production = 6
		r.NextUpgradeTime = 90
		r.NextUpgradeCost["iron"] = 0
		r.NextUpgradeCost["copper"] = 800
		r.NextUpgradeCost["gold"] = 16
	case 5:
		r.Production = 8
		r.NextUpgradeTime = 120
	}

	r.Level = level
}

func (r *GoldFactory) Upgrade() {
	if r.Level == 5 || r.IsUpgrade {
		return
	}

	r.IsUpgrade = true
	time.Sleep(r.NextUpgradeTime)

	r.Level += 1
	r.UpgradeToLevel(r.Level)
	r.updCb()

	r.IsUpgrade = false
}

func newGoldFactory(upd func()) GoldFactory {
	resource := GoldFactory{
		Kind:  "gold",
		updCb: upd,
	}
	resource.NextUpgradeCost = make(map[string]int)

	resource.Upgrade()

	return resource
}
