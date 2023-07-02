package model

type UserType struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
