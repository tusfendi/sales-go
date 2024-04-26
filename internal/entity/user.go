package entity

import "time"

type User struct {
	ID        int64      `gorm:"column:id; PRIMARY KEY"`
	Name      string     `gorm:"column:name" json:"name" binding:"required"`
	Email     string     `gorm:"column:email" json:"username" binding:"required"`
	Password  string     `gorm:"column:password" json:"password" binding:"required"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
