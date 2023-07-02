package model

type UserPosition struct {
	ID   int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}
