package response

import (
	"hr-management-system/model"
	"strings"
)

type LoginReponse struct {
	ID                 int                 `json:"id"`
	Email              string              `json:"email"`
	CompanyResponse    CompanyResponse     `json:"company"`
	UserProfileReponse UserProfileResponse `json:"user"`
	Roles              []string            `json:"roles"`
	Token              string              `json:"token"`
}

func ResponseLogin(user model.User, token string) LoginReponse {
	user.Roles = strings.Trim(user.Roles, "{}")
	roles := strings.Split(user.Roles, ",")
	return LoginReponse{
		ID:              user.ID,
		Email:           user.Email,
		CompanyResponse: ResponseCompany(user.Company),
		UserProfileReponse: UserProfileResponse{
			Fullname:         user.UserProfile.FullName,
			ProfilePicture:   user.UserProfile.ProfilePicture,
			TypeUserCode:     user.UserProfile.TypeID,
			PositionUserCode: user.UserProfile.PositionID,
			PositionUserName: user.UserProfile.UserPosition.Name,
		},
		Roles: roles,
		Token: token,
	}
}
