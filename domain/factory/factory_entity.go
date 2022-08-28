package factory

import (
	"time"
)

type factoryUpgradeMap struct {
	production      int
	nextUpgradeTime time.Duration
	nextUpgradeCost map[string]int
}

type Factory struct {
	IsRunning       bool           `json:"isRunning"`
	Kind            string         `json:"kind"`
	Total           int            `json:"total"`
	Level           int            `json:"level"`
	IsUpgrade       bool           `json:"isUpgrade"`
	Production      int            `json:"productionTime"`
	NextUpgradeTime time.Duration  `json:"nextUpgradeTime"`
	NextUpgradeCost map[string]int `json:"nextUpgradeCost"`
	upgradeMap      map[int]factoryUpgradeMap
	updCb           func(total int, level int)
}

func (r *Factory) Start() {
	r.IsRunning = true

	if r.Level == 0 {
		r.UpgradeToLevel(1)
	} else {
		r.UpgradeToLevel(r.Level)
	}

	r.Loop()
}

func (r *Factory) Stop() {
	r.IsRunning = false
}

func (r *Factory) Loop() {
	if !r.IsRunning {
		return
	}

	time.Sleep(time.Second)
	r.Total += r.Production
	r.Loop()
}

func (f *Factory) UpgradeToLevel(level int) {
	upgradeMap, ok := f.upgradeMap[level]
	if !ok {
		return
	}

	f.Production = upgradeMap.production
	f.NextUpgradeTime = upgradeMap.nextUpgradeTime

	for key, val := range upgradeMap.nextUpgradeCost {
		f.NextUpgradeCost[key] = val
	}

	f.Level = level
}

func (r *Factory) Upgrade() {
	if r.Level == 5 || r.IsUpgrade {
		return
	}

	r.IsUpgrade = true
	time.Sleep(r.NextUpgradeTime)

	r.Level += 1
	r.UpgradeToLevel(r.Level)
	r.updCb(r.Total, r.Level)

	r.IsUpgrade = false
}

func newFactory(kind string, total int, level int, upgradeMap map[int]factoryUpgradeMap) *Factory {
	resource := Factory{
		Kind:            kind,
		Total:           total,
		Level:           level,
		NextUpgradeCost: make(map[string]int),
		upgradeMap:      upgradeMap,
	}

	return &resource
}
