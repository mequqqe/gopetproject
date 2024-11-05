package database

import (
	"goprometheus/internal/domain/models"
	"log"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}).Error; err != nil {
		return err
	}
	log.Println("Migrations applied successfully!")
	return nil
}
