// /internal/app/usecase/create_user.go
package usecase

import (
	"fmt"
	"goprometheus/internal/domain/models"
	"goprometheus/internal/interface/repository"
)

type CreateUserUseCase struct {
	UserRepo repository.UserRepositoryImpl
}

// Execute - Логика для создания пользователя
func (uc *CreateUserUseCase) Execute(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return fmt.Errorf("invalid user data: name and email are required")
	}

	// Используем репозиторий для сохранения пользователя
	err := uc.UserRepo.Save(user)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}

	return nil
}
