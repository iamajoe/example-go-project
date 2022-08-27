package sqlite

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

func Connect(file string) (*DB, error) {
	// create the file if doesnt exist
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil

	// defer db.Close()

	// var version string
	// err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	// if err != nil {
	// 		log.Fatal(err)
	// }

	// return nil, nil
}
