package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joesantosio/example-go-project/config"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/infrastructure/inmem"
	"github.com/joesantosio/example-go-project/infrastructure/sqlite"
	"github.com/joesantosio/example-go-project/interfaces/server"
	serverchi "github.com/joesantosio/example-go-project/interfaces/server_chi"
	serverstd "github.com/joesantosio/example-go-project/interfaces/server_std"
)

func initRepos(dbType string, dbPath string) (repos entity.Repositories, err error) {
	switch dbType {
	case "sqlite":
		if dbPath == "" {
			err = errors.New("DB_PATH not provided")
			break
		}

		db, err := sqlite.Connect(dbPath)
		if err != nil {
			err = errors.New(fmt.Sprintf("error initializing db sqlite: %v", err))
			break
		}

		repos, err = sqlite.InitRepos(db)
		if err != nil {
			err = errors.New(fmt.Sprintf("error initializing repos sqlite: %v", err))
			break
		}
	default:
		repos, err = inmem.InitRepos()
		if err != nil {
			err = errors.New(fmt.Sprintf("error initializing repos inmem: %v", err))
			break
		}
	}

	return repos, nil
}

func main() {
	config, err := config.Get(os.Getenv)
	if err != nil {
		log.Fatal(err)
		return
	}

	repos, err := initRepos(config.DBType, config.DBPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer repos.Close()

	var r http.Handler

	switch os.Getenv("SERVER_PACKAGE") {
	case "chi":
		r = serverchi.GetRouter(config.AuthSecret, repos)
		break
	default:
		r = serverstd.GetRouter(config.AuthSecret, repos)
	}

	server.InitServer(":4040", r)
}
