package repository

import (
	"hr-management-system/app/entity"

	"gorm.io/gorm"
)

// Interface
type UserRepository interface {
	FindByEmail(email string) (entity.User, error)
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
func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("users.email = ?", email).Where("users.active = ?", true).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
