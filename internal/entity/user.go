package entity

import "time"

type User struct {
	ID        int64      `gorm:"column:id; PRIMARY KEY"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	Password  string     `gorm:"column:password"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
