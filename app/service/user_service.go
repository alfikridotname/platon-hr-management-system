package service

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type service interface {
	Login(input LoginInput) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}
