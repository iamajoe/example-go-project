package main

import (
	"fmt"
	"net/http"
)

func main() {
	repos, err := initRepos(false)
	if err != nil {
		panic(err)
	}
	defer repos.Close()

	http.HandleFunc("/user", reqCreateUser(repos))
	http.HandleFunc("/dashboard", reqGetDashboard(repos))
	http.HandleFunc("/factory/upgrade", reqUpgradeFactory(repos))

	fmt.Println("listening at :4040")
	err = http.ListenAndServe(":4040", nil)
	if err != nil {
		panic(err)
	}
}
