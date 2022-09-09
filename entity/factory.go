package entity

type Factory struct {
	ID     int
	UserID string
	Kind   string `json:"kind"`
	Total  int    `json:"total"`
	Level  int    `json:"level"`
}

func NewModelFactory(id int, userID string, kind string, total int, level int) *Factory {
	return &Factory{
		ID:     id,
		UserID: userID,
		Kind:   kind,
		Total:  total,
		Level:  level,
	}
}

type RepositoryFactory interface {
	GetByUserID(userID string) ([]Factory, error)
	Create(kind string, total int, level int, userID string) (bool, error)
	Patch(kind string, userID string, total int, level int) (bool, error)
	RemoveFactoriesFromUser(userID string) (bool, error)
	Remove(kind string, userID string) (bool, error)
}
