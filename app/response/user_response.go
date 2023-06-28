package response

import "hr-management-system/app/entity"

type UserReponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ResponseUser(user entity.User, token string) UserReponse {
	return UserReponse{
		ID:    user.ID,
		Email: user.Email,
		Token: token,
	}
}
