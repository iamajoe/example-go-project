package infrastructure

type Repositories struct {
	closeFn func() error
	User    RepositoryUser
	Factory RepositoryFactory
}

func (r *Repositories) Close() error {
	if r.closeFn != nil {
		return r.closeFn()
	}

	return nil
}

func NewRepositories(user RepositoryUser, factory RepositoryFactory, closeFn func() error) *Repositories {
	return &Repositories{closeFn, user, factory}
}
