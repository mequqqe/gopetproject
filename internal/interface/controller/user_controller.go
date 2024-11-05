// /internal/interface/controller/user_controller.go
package controller

import (
	"encoding/json"
	"goprometheus/internal/app/usecase"
	"goprometheus/internal/domain/models"
	"net/http"
)

// UserController - Контроллер для работы с пользователями
type UserController struct {
	CreateUserUseCase *usecase.CreateUserUseCase
	GetUsersUseCase   *usecase.GetUsersUseCase
}

// CreateUser - Обработчик для создания пользователя
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Создаем нового пользователя
	user := &models.User{Name: request.Name, Email: request.Email}
	err := uc.CreateUserUseCase.Execute(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ с созданным пользователем
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.CreateUserUseCase.UserRepo.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
