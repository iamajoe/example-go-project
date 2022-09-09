package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/infrastructure/inmem"
	"github.com/joesantosio/simple-game-api/infrastructure/sqlite"
	"github.com/joesantosio/simple-game-api/interfaces/http_chi"
	"github.com/joesantosio/simple-game-api/interfaces/http_std"
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
	repos, err := initRepos(os.Getenv("DB_TYPE"), os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer repos.Close()

	authSecret := os.Getenv("AUTH_SECRET")
	if authSecret == "" {
		authSecret = fmt.Sprintf("%d", rand.Intn(100000000))
	}

	switch os.Getenv("SERVER_PACKAGE") {
	case "chi":
		http_chi.InitServer(":4040", "secret", repos)
		break
	default:
		http_std.InitServer(":4040", repos)
	}
}
