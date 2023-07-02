package repository

import (
	"hr-management-system/model"

	"gorm.io/gorm"
)

// Interface
type UserRepository interface {
	FindByEmail(email string) (model.User, error)
}

// Struct
type userRepository struct {
	db *gorm.DB
}

// Constructor
func NewRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Method
func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Preload("Company").
		Preload("Company.CompanyStatus").
		Preload("UserStatus").
		Preload("UserProfile").
		Preload("UserProfile.UserType").
		Preload("UserProfile.UserPosition").
		Where("users.email = ?", email).
		Find(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
