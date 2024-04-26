package entity

import "time"

type Sales struct {
	ID              int64      `gorm:"column:id; PRIMARY KEY"`
	UserID          int64      `gorm:"column:user_id"`
	Type            string     `gorm:"column:type"`
	Nominal         int64      `gorm:"column:nominal"`
	TransactionDate time.Time  `gorm:"column:transaction_date"`
	CreatedAt       time.Time  `gorm:"created_at"`
	UpdatedAt       *time.Time `gorm:"updated_at"`
}

type DownlaodSalesData struct {
	Name                  string `gorm:"column:name"`
	DaysTotal             int64  `gorm:"column:days_total"`
	TransactionTotal      int64  `gorm:"column:transaction_total"`
	GoodsTransactionTotal int64  `gorm:"column:goods_transaction_total"`
	NominalTotal          int64  `gorm:"column:nominal_total"`
	GoodsNominalTotal     int64  `gorm:"column:goods_nominal_total"`
}

func (Sales) TableName() string {
	return "sales"
}
