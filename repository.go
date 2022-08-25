package main

type Repositories struct {
	db      *DBSqlite
	user    RepositoryUser
	factory RepositoryFactory
}

func (r *Repositories) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

func initDBRepos(db *DBSqlite) (Repositories, error) {
	if db == nil {
		return Repositories{}, nil
	}

	userRepo, err := createRepositoryUserSqlite(db)
	if err != nil {
		return Repositories{}, err
	}

	factoryRepo, err := createRepositoryFactorySqlite(db)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{db, userRepo, factoryRepo}, nil
}

func initInmemRepos() (Repositories, error) {
	factoryRepo, err := createRepositoryFactoryInmem()
	if err != nil {
		return Repositories{}, err
	}

	userRepo, err := createRepositoryUserInmem(factoryRepo)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{nil, userRepo, factoryRepo}, nil
}

func initRepos(isInmem bool) (Repositories, error) {
	if isInmem {
		return initInmemRepos()
	}

	db, err := ConnectSQLite("data.db")
	if err != nil {
		return Repositories{}, err
	}

	return initDBRepos(db)
}
