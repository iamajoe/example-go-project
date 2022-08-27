package main

import (
	"log"
	"os"
	"strings"

	"github.com/joesantosio/simple-game-api/infrastructure/inmem"
	"github.com/joesantosio/simple-game-api/infrastructure/sqlite"
	"github.com/joesantosio/simple-game-api/interfaces/httpd"
)

func main() {
	repos, err := inmem.InitRepos()
	if err != nil {
		log.Fatalf("error initializing repos inmem: %v", err)
		return
	}

	// check if tests, if not, we can setup the sqlite
	if !strings.Contains(os.Args[0], "/_test/") {
		db, err := sqlite.Connect("data.sqlite")
		if err != nil {
			log.Fatalf("error initializing db sqlite: %v", err)
			return
		}

		repos, err = sqlite.InitRepos(db)
		if err != nil {
			log.Fatalf("error initializing repos sqlite: %v", err)
			return
		}
	}

	defer repos.Close()

	httpd.InitServer(":4040", repos)
}
