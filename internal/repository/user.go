package repository

import (
	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/entity"
)

type UserRepository interface {
	TrxSupportRepo
	GetByEmail(email string) (entity.User, bool, error)
	CreateUser(user *entity.User) (result *entity.User, err error)
}

type userRepoCtx struct {
	GormTrxSupport
}

func NewUserRepository(db *config.Mysql) *userRepoCtx {
	return &userRepoCtx{
		GormTrxSupport{db: db.DB},
	}
}

func (r *userRepoCtx) GetByEmail(email string) (entity.User, bool, error) {
	user := entity.User{}
	res := r.db.Where("email = ?", email).Find(&user)
	if res.RowsAffected == 0 || res.Error != nil {
		return entity.User{}, false, res.Error
	}
	return user, res.RowsAffected > 0, nil
}

func (r *userRepoCtx) CreateUser(user *entity.User) (result *entity.User, err error) {
	err = r.db.Create(user).Error
	return user, err
}
