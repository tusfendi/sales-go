package repository

import (
	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/entity"
)

type SalesRepository interface {
	CreateSales(sales *entity.Sales) error
	TrxSupportRepo
}

type salesRepoCtx struct {
	GormTrxSupport
}

func NewSalesRepository(db *config.Mysql) *salesRepoCtx {
	return &salesRepoCtx{
		GormTrxSupport{db: db.DB},
	}
}

func (r *salesRepoCtx) CreateSales(sales *entity.Sales) error {
	return r.db.Create(sales).Error
}
