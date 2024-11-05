package main

import (
	"goprometheus/internal/app/usecase"
	"goprometheus/internal/database"
	"goprometheus/internal/interface/controller"
	"goprometheus/internal/interface/repository"
	"log"
	"net/http"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	userRepo := &repository.UserRepositoryImpl{DB: db}

	// Создание сценариев
	createUserUseCase := &usecase.CreateUserUseCase{UserRepo: *userRepo}
	getUsersUseCase := &usecase.GetUsersUseCase{UserRepo: *userRepo}

	// Создание контроллера
	userController := &controller.UserController{
		CreateUserUseCase: createUserUseCase,
		GetUsersUseCase:   getUsersUseCase, // Передаем GetUsersUseCase
	}

	// Настройка маршрутов
	http.HandleFunc("/users", userController.CreateUser)
	http.HandleFunc("/users/list", userController.GetUsers)

	// Запуск HTTP-сервера
	log.Fatal(http.ListenAndServe(":8080", nil))
}
