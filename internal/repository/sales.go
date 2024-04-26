package repository

import (
	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/entity"
	"github.com/tusfendi/sales-go/internal/presenter"
)

type SalesRepository interface {
	CreateSales(sales *entity.Sales) error
	GetData(params presenter.DownloadSalesRequest) ([]entity.DownlaodSalesData, error)
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

func (r *salesRepoCtx) GetData(params presenter.DownloadSalesRequest) ([]entity.DownlaodSalesData, error) {
	var result []entity.DownlaodSalesData
	q := r.db.Raw(`SELECT u.name, data.* FROM users u 
		LEFT JOIN 
			(SELECT SUM(nominal) as nominal_total, user_id, count(*) as transaction_total,
			SUM(CASE WHEN type = 'goods' THEN nominal ELSE 0 END) goods_nominal_total,
			SUM(CASE WHEN type = 'goods' THEN 1 ELSE 0 END) goods_transaction_total,
			count(DISTINCT(transaction_date)) as days_total
			FROM sales s WHERE s.transaction_date >= ? AND s.transaction_date <= ? GROUP BY user_id) 
		as data ON data.user_id = u.id;`, params.StartDate, params.EndDate).Find(&result)

	return result, q.Error
}
