package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func reqUpgradeFactory(userRepo RepositoryUser, factoryRepo RepositoryFactory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		user, err := getUserByUsername(username, userRepo, factoryRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var bodyResource map[string]string
		if err = json.NewDecoder(r.Body).Decode(&bodyResource); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ok, err := user.UpgradeResource(bodyResource["kind"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		io.WriteString(w, fmt.Sprintf("%t", ok))
	}
}

func reqCreateUser(userRepo RepositoryUser, factoryRepo RepositoryFactory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ok, err := createUser(user.Username, userRepo, factoryRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		io.WriteString(w, fmt.Sprintf("%t", ok))
	}
}

func reqGetDashboard(userRepo RepositoryUser, factoryRepo RepositoryFactory) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		user, err := getUserByUsername(username, userRepo, factoryRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func main() {
	db, err := ConnectSQLite("data.db")
	if err != nil {
		// TODO: could probably handle it more gracefully
		panic(err)
	}
	// TODO: should handle the database with an interface
	defer db.db.Close()

	userRepo, err := createRepositoryUserSqlite(db)
	if err != nil {
		// TODO: could probably handle it more gracefully
		panic(err)
	}

	factoryRepo, err := createRepositoryFactorySqlite(db)
	if err != nil {
		// TODO: could probably handle it more gracefully
		panic(err)
	}

	http.HandleFunc("/user", reqCreateUser(userRepo, factoryRepo))
	http.HandleFunc("/dashboard", reqGetDashboard(userRepo, factoryRepo))
	http.HandleFunc("/upgrade", reqUpgradeFactory(userRepo, factoryRepo))

	fmt.Println("listening at :4040")
	err = http.ListenAndServe(":4040", nil)
	if err != nil {
		// TODO: could probably handle it more gracefully
		panic(err)
	}
}
