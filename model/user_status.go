package model

type UserStatus struct {
	ID   int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

func (UserStatus) TableName() string {
	return "user_status"
}
