// /internal/database/db.go
package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Импортируем драйвер PostgreSQL
)

// InitDB - Функция для подключения к базе данных PostgreSQL с помощью GORM
func InitDB() (*gorm.DB, error) {
	// Строка подключения (DSN)
	dsn := "user=postgres password=1234 dbname=petproject sslmode=disable"

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	db.LogMode(true) // Включить вывод SQL-запросов для отладки (по желанию)
	return db, nil
}
