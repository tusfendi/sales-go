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

func (Sales) TableName() string {
	return "sales"
}
