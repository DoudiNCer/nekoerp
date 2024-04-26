package tables

import "time"

type Tiao struct {
	Id        uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Type      uint      `gorm:"column:type"`
	Count     uint      `gorm:"column:count"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Operator  uint      `gorm:"column:operator"`
}

func (tiao Tiao) TableName() string {
	return "tiao"
}
