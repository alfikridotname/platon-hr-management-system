package response

import "hr-management-system/model"

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

func ResponseUserProfile(userProfile model.UserProfile, user model.User) UserProfileResponse {
	return UserProfileResponse{
		StatusUserCode:   user.StatusID,
		StatusUserName:   user.UserStatus.Name,
		Fullname:         userProfile.FullName,
		ProfilePicture:   userProfile.ProfilePicture,
		TypeUserCode:     userProfile.TypeID,
		TypeUserName:     userProfile.UserType.Name,
		PositionUserCode: userProfile.PositionID,
		PositionUserName: userProfile.UserPosition.Name,
	}
}
