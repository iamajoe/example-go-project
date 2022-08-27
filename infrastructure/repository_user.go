package infrastructure

type User interface {
	GetUsername() string
	SetUsername(u string)
}

type RepositoryUser interface {
	GetUserByUsername(username string) (User, error)
	CreateUser(user User) (bool, error)
}

type modelUser struct {
	Username string `json:"username"`
}

func (u *modelUser) GetUsername() string {
	return u.Username
}
func (u *modelUser) SetUsername(username string) {
	u.Username = username
}

func NewUser(username string) User {
	return &modelUser{username}
}
