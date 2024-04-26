package server

import (
	"sync"

	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/repository"
)

type RepositoryRegistry interface {
	UserRepo() repository.UserRepository
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
