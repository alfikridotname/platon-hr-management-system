package response

type UserProfileResponse struct {
	StatusUserCode   int    `json:"status_user_code"`
	StatusUserName   string `json:"status_user_name"`
	Fullname         string `json:"fullname"`
	ProfilePicture   string `json:"profile_picture"`
	TypeUserCode     int    `json:"type_user_code"`
	TypeUserName     string `json:"type_user_name"`
	PositionUserCode int    `json:"position_user_code"`
	PositionUserName string `json:"position_user_name"`
}
