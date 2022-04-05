package models

// Tax structure
type Tax struct {
	Id     int    `gorm:"primaryKey;type:int;not null" json:"id"`
	Name   string `gorm:"type:varchar(450)" json:"name"`
	Rate   int    `gorm:"type:int" json:"rate"`
	ItemId int    `gorm:"column:item_id"`
}

// TableName return name of database table
func (t *Tax) TableName() string {
	return "pajak"
}
