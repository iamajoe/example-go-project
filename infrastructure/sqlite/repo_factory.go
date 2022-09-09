package sqlite

import (
	"github.com/joesantosio/example-go-project/entity"
)

// ------------------------------
// repository

type repositoryFactory struct {
	tableName string
	db        *DB
}

func (repo *repositoryFactory) GetByUserID(userID string) ([]entity.Factory, error) {
	factories := []entity.Factory{}

	rows, err := repo.db.db.Query("SELECT id,userid,kind,total,level FROM "+repo.tableName+" WHERE userid=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var userID string
		var kind string
		var total int
		var level int

		err = rows.Scan(&id, &userID, &kind, &total, &level)
		if err != nil {
			return nil, err
		}

		factory := entity.NewModelFactory(id, userID, kind, total, level)
		factories = append(factories, *factory)
	}

	return factories, nil
}

func (repo *repositoryFactory) Create(kind string, total int, level int, userID string) (bool, error) {
	sts := "INSERT INTO " + repo.tableName + "(userid, kind, total, level) VALUES(?, ?, ?, ?)"
	_, err := repo.db.db.Exec(sts, userID, kind, total, level)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryFactory) Patch(kind string, userID string, total int, level int) (bool, error) {
	sts := "UPDATE " + repo.tableName + " SET total=?, level=? WHERE userid=? AND kind=?"
	_, err := repo.db.db.Exec(sts, total, level, userID, kind)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repo *repositoryFactory) RemoveFactoriesFromUser(userID string) (bool, error) {
	sts := "DELETE FROM " + repo.tableName + " WHERE userid=?"
	_, err := repo.db.db.Exec(sts, userID)
	return true, err
}

func (repo *repositoryFactory) Remove(kind string, userID string) (bool, error) {
	sts := "DELETE FROM " + repo.tableName + " WHERE userid=? AND kind=?"
	_, err := repo.db.db.Exec(sts, userID, kind)
	return true, err
}

func (repo *repositoryFactory) RemoveTable() (bool, error) {
	sts := "DROP TABLE IF EXISTS " + repo.tableName + " CASCADE"
	_, err := repo.db.db.Exec(sts)
	return true, err
}

func createRepositoryFactory(db *DB) (entity.RepositoryFactory, error) {
	repo := repositoryFactory{"factories", db}

	// TODO: should have a reference to the users table
	sts := `
		CREATE TABLE IF NOT EXISTS factories(
			id              INTEGER PRIMARY KEY AUTOINCREMENT,
			userid 					TEXT, 
			kind 						TEXT, 
			total 					INTEGER, 
			level 					INTEGER
		);
	`
	_, err := repo.db.db.Exec(sts)

	return &repo, err
}
