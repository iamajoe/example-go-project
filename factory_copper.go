package main

import "time"

type CopperFactory struct {
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

func (r *CopperFactory) GetKind() string {
	return r.Kind
}

func (r *CopperFactory) GetTotal() int {
	return r.Total
}

func (r *CopperFactory) GetLevel() int {
	return r.Level
}

func (r *CopperFactory) GetNextUpgradeCost() map[string]int {
	return r.NextUpgradeCost
}

func (r *CopperFactory) Start() {
	r.IsRunning = true

	if r.Level == 0 {
		r.UpgradeToLevel(1)
	} else {
		r.UpgradeToLevel(r.Level)
	}

	r.Loop()
}

func (r *CopperFactory) Stop() {
	r.IsRunning = false
}

func (r *CopperFactory) Loop() {
	if !r.IsRunning {
		return
	}

	time.Sleep(time.Second)
	r.Total += r.Production
	r.updCb()
	r.Loop()
}

func (r *CopperFactory) UpgradeToLevel(level int) {
	switch level {
	case 1:
		r.Production = 10
		r.NextUpgradeTime = 15
		r.NextUpgradeCost["iron"] = 300
		r.NextUpgradeCost["copper"] = 100
		r.NextUpgradeCost["gold"] = 1
	case 2:
		r.Production = 20
		r.NextUpgradeTime = 30
		r.NextUpgradeCost["iron"] = 800
		r.NextUpgradeCost["copper"] = 250
		r.NextUpgradeCost["gold"] = 2
	case 3:
		r.Production = 40
		r.NextUpgradeTime = 60
		r.NextUpgradeCost["iron"] = 1600
		r.NextUpgradeCost["copper"] = 500
		r.NextUpgradeCost["gold"] = 4
	case 4:
		r.Production = 80
		r.NextUpgradeTime = 90
		r.NextUpgradeCost["iron"] = 3000
		r.NextUpgradeCost["copper"] = 1000
		r.NextUpgradeCost["gold"] = 8
	case 5:
		r.Production = 150
		r.NextUpgradeTime = 120
	}

	r.Level = level
}

func (r *CopperFactory) Upgrade() {
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

func newCopperFactory(upd func()) CopperFactory {
	resource := CopperFactory{
		Kind:  "copper",
		updCb: upd,
	}
	resource.NextUpgradeCost = make(map[string]int)

	resource.Upgrade()

	return resource
}
