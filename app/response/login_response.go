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
		ID:                 user.ID,
		Email:              user.Email,
		CompanyResponse:    ResponseCompany(user.Company),
		UserProfileReponse: ResponseUserProfile(user.UserProfile, user),
		Roles:              roles,
		Token:              token,
	}
}
