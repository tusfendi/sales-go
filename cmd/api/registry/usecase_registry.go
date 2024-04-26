package registry

import (
	"sync"

	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/usecase"
)

type UsecaseRegistry interface {
	User() usecase.UserUsecase
	Sales() usecase.SalesUsecase
}

type usecaseRegistry struct {
	cfg  *config.Config
	repo RepositoryRegistry
}

func NewUsecaseRegistry(cfg *config.Config, repo RepositoryRegistry) (r UsecaseRegistry) {
	var ucRegistry usecaseRegistry
	var once sync.Once

	once.Do(func() {
		ucRegistry = usecaseRegistry{
			cfg:  cfg,
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
			*u.cfg,
			u.repo.UserRepo(),
		)
	})
	return userUsecase
}

func (u *usecaseRegistry) Sales() usecase.SalesUsecase {
	var salesUsecase usecase.SalesUsecase
	var once sync.Once
	once.Do(func() {
		salesUsecase = usecase.NewSalesUsecase(
			u.repo.SalesRepo(),
		)
	})
	return salesUsecase
}
