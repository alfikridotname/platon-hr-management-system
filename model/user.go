package model

type User struct {
	ID          int         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Email       string      `gorm:"column:email;unique;not null" json:"email"`
	Password    string      `gorm:"column:password;not null" json:"password"`
	CompanyID   int         `gorm:"column:company_id;not null" json:"company_id"`
	Company     Company     `gorm:"foreignkey:CompanyID" json:"company"`
	StatusID    int         `gorm:"column:status_id;not null" json:"status_id"`
	UserStatus  UserStatus  `gorm:"foreignkey:StatusID" json:"user_status"`
	UserProfile UserProfile `gorm:"foreignkey:ID" json:"user_profile"`
	Roles       string      `gorm:"column:roles;not null" json:"roles"`
}
