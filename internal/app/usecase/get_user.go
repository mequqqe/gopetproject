package usecase

import (
	"fmt"
	"goprometheus/internal/domain/models"
	"goprometheus/internal/interface/repository"
)

type GetUsersUseCase struct {
	UserRepo repository.UserRepositoryImpl
}

func (uc *GetUsersUseCase) Execute() ([]models.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("error при получении юзеров", err)
	}
	return users, nil
}
