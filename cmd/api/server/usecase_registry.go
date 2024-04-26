package server

import (
	"sync"

	"github.com/tusfendi/sales-go/internal/usecase"
)

type UsecaseRegistry interface {
	User() usecase.UserUsecase
}

type usecaseRegistry struct {
	repo RepositoryRegistry
}

func NewUsecaseRegistry(repo RepositoryRegistry) (r UsecaseRegistry) {
	var ucRegistry usecaseRegistry
	var once sync.Once

	once.Do(func() {
		ucRegistry = usecaseRegistry{
			repo: repo,
		}
	})

	return &ucRegistry
}

func (u *usecaseRegistry) User() usecase.UserUsecase {
	var userUsecase usecase.UserUsecase
	var once sync.Once
	once.Do(func() {
		userUsecase = usecase.NewUserUsecase(
			u.repo.UserRepo(),
		)
	})
	return userUsecase
}
