package main

type Factory interface {
	GetKind() string
	GetTotal() int
	GetLevel() int
	GetNextUpgradeCost() map[string]int
	Start()
	Upgrade()
}
