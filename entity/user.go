package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewModelUser(id string, username string) User {
	return User{id, username}
}

type RepositoryUser interface {
	GetByIDs(ids []string) ([]User, error)
	GetByUsername(username string) (User, error)
	Create(username string) (string, error)
}

type UserToken struct {
	ID     int
	UserId string
	Token  string
}

func NewModelUserToken(id int, userId string, token string) UserToken {
	return UserToken{id, userId, token}
}

type RepositoryUserToken interface {
	Create(userId string, token string) (bool, error)
	IsTokenValid(token string) (bool, error)
}
