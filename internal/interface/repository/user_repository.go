// /internal/interface/repository/user_repository.go
package repository

import (
	"fmt"
	"goprometheus/internal/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *UserRepositoryImpl) Save(user *models.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
	return nil
}

func (repo *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}
	return users, nil
}
