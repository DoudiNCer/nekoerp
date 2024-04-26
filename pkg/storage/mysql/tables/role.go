package tables

type Role struct {
	Id   uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name string `gorm:"column:name"`
}

func (role Role) TableName() string {
	return "role"
}
