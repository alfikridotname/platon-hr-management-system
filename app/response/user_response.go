package response

import "hr-management-system/model"

type UserReponse struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CompanyID int    `json:"company_id"`
}

func ResponseUser(user model.User) UserReponse {
	return UserReponse{
		ID:        user.ID,
		Email:     user.Email,
		CompanyID: user.CompanyID,
	}
}
