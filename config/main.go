package config

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Config struct {
	AuthSecret    string
	Port          string
	DBType        string
	DBPath        string
	ServerPackage string
}

func Get(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}

	authSecret := getenv("AUTH_SECRET")
	if authSecret == "" {
		rand.Seed(time.Now().Unix())
		authSecret = fmt.Sprintf("%d", rand.Intn(100000000))
	}

	return Config{
		AuthSecret:    authSecret,
		Port:          getenv("PORT"),
		DBType:        getenv("DB_TYPE"),
		DBPath:        getenv("DB_PATH"),
		ServerPackage: getenv("SERVER_PACKAGE"),
	}, nil
}
