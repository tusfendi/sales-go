package usecase

import "github.com/tusfendi/sales-go/internal/repository"

type UserUsecase interface {
}

type userCtx struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userCtx{
		userRepo: userRepo,
	}
}
