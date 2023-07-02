package model

type UserProfile struct {
	ID             int          `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserID         int          `gorm:"column:user_id;not null" json:"user_id"`
	FullName       string       `gorm:"column:fullname;not null" json:"fullname"`
	ProfilePicture string       `gorm:"column:profile_picture;not null" json:"profile_picture"`
	TypeID         int          `gorm:"column:type_id;not null" json:"type_id"`
	PositionID     int          `gorm:"column:position_id;not null" json:"position_id"`
	UserPosition   UserPosition `gorm:"foreignkey:PositionID" json:"user_position"`
}
