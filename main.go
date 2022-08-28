package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/infrastructure/inmem"
	"github.com/joesantosio/simple-game-api/infrastructure/sqlite"
	"github.com/joesantosio/simple-game-api/interfaces/http_std"
)

func initRepos() (entity.Repositories, error) {
	repos, err := inmem.InitRepos()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error initializing repos inmem: %v", err))
	}

	// check if tests, if not, we can setup the sqlite
	if !strings.Contains(os.Args[0], "/_test/") {
		db, err := sqlite.Connect("data.sqlite")
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error initializing db sqlite: %v", err))
		}

		repos, err = sqlite.InitRepos(db)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error initializing repos sqlite: %v", err))
		}
	}

	return repos, nil
}

func main() {
	repos, err := initRepos()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer repos.Close()

	switch os.Getenv("SERVER_PACKAGE") {
	default:
		http_std.InitServer(":4040", repos)
	}
}
