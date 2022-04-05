package models

// Item structure
type Item struct {
	Id    int    `gorm:"primaryKey;type:int;not null" json:"id"`
	Name  string `gorm:"type:varchar(450)" json:"name"`
	Taxes []Tax  `gorm:"ForeignKey:ItemId" json:"taxes"`
}

// TableName return name of database table
func (t *Item) TableName() string {
	return "item"
}
