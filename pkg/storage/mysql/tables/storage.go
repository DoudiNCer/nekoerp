package tables

type Storage struct {
	Id    uint `gorm:"primaryKey;column:id"`
	Type  uint `gorm:"column:type"`
	Count uint `gorm:"column:count"`
}

func (storage Storage) TableName() string {
	return "storage"
}
