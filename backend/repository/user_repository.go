package repository

import (
	"fmt"

	"github.com/tigaweb/reversi-app/backend/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(&user).Error; err != nil {
		return fmt.Errorf("incorrect email address or password")
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
