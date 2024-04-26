package tables

type Goods struct {
	Id    uint    `gorm:"primaryKey;column:id"`
	Name  string  `gorm:"column:name"`
	Price float64 `gorm:"type:numeric(10,2);column:price"`
}

func (goods Goods) TableName() string {
	return "goods"
}
