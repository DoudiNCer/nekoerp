package tables

import "time"

type Di struct {
	Id        uint      `gorm:"primaryKey;column:id"`
	Type      uint      `gorm:"column:type"`
	Count     uint      `gorm:"column:count"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Operator  uint      `gorm:"column:operator"`
}

func (di Di) TableName() string {
	return "di"
}
