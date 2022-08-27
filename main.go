package main

func main() {
	repos, err := initRepos(false)
	if err != nil {
		panic(err)
	}
	defer repos.Close()

	initServer(repos)
}
