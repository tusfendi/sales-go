package registry

import (
	"sync"

	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/repository"
)

type RepositoryRegistry interface {
	UserRepo() repository.UserRepository
	SalesRepo() repository.SalesRepository
}

type reposiotryRegistry struct {
	cfg *config.Config
	db  *config.Mysql
}

func NewRepositoryRegistry(cfg *config.Config, msqlConn *config.Mysql) (r RepositoryRegistry) {
	var repoRegistry reposiotryRegistry
	var once sync.Once

	once.Do(func() {
		repoRegistry = reposiotryRegistry{
			cfg: cfg,
			db:  msqlConn,
		}
	})

	return &repoRegistry
}

func (r reposiotryRegistry) UserRepo() repository.UserRepository {
	var once sync.Once
	var userRepo repository.UserRepository

	once.Do(func() {
		userRepo = repository.NewUserRepository(r.db)
	})

	return userRepo
}

func (r reposiotryRegistry) SalesRepo() repository.SalesRepository {
	var once sync.Once
	var salesRepo repository.SalesRepository

	once.Do(func() {
		salesRepo = repository.NewSalesRepository(r.db)
	})

	return salesRepo
}
