package service

import (
	"errors"
	"hr-management-system/app/entity"
	"hr-management-system/app/repository"
	"hr-management-system/app/request"

	"golang.org/x/crypto/bcrypt"
)

// Interface
type UserService interface {
	Login(input request.LoginInput) (entity.User, error)
}

// Struct
type userService struct {
	repository repository.UserRepository
}

// Constructor
func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

// Method
func (s *userService) Login(input request.LoginInput) (entity.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("USER NOT FOUND")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
