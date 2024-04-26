package repository

import "github.com/tusfendi/sales-go/config"

type UserRepository interface {
	TrxSupportRepo
}

type userRepoCtx struct {
	GormTrxSupport
}

func NewUserRepository(db *config.Mysql) *userRepoCtx {
	return &userRepoCtx{
		GormTrxSupport{db: db.DB},
	}
}
