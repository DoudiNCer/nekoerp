package tables

import "time"

type User struct {
	Id        uint      `gorm:"primaryKey;column:id"`
	Account   string    `gorm:"column:account"`
	Password  string    `gorm:"column:password"`
	Role      uint      `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Blocked   bool      `gorm:"column:blocked"`
}

func (user User) TableName() string {
	return "user"
}
