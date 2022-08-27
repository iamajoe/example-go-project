package infrastructure

// type Factory interface {
// 	GetKind() string
// 	GetTotal() int
// 	GetLevel() int
// 	GetNextUpgradeCost() map[string]int
// 	Start()
// 	Upgrade()
// }

type Factory interface {
	GetKind() string
	GetTotal() int
	SetTotal(t int)
	GetLevel() int
	SetLevel(l int)
}

type RepositoryFactory interface {
	GetByUsername(username string) ([]Factory, error)
	CreateFactory(factory Factory, username string) (bool, error)
	PatchFactory(factory Factory, username string) (bool, error)
	RemoveFactoriesFromUser(username string) (bool, error)
	RemoveFactory(factory Factory, username string) (bool, error)
}

type modelFactory struct {
	Kind  string `json:"kind"`
	Total int    `json:"total"`
	Level int    `json:"level"`
}

func (r *modelFactory) GetKind() string {
	return r.Kind
}

func (r *modelFactory) GetTotal() int {
	return r.Total
}
func (r *modelFactory) SetTotal(t int) {
	r.Total = t
}

func (r *modelFactory) GetLevel() int {
	return r.Level
}
func (r *modelFactory) SetLevel(l int) {
	r.Level = l
}

func NewFactory(kind string, total int, level int) Factory {
	return &modelFactory{
		Kind:  kind,
		Total: total,
		Level: level,
	}
}
