package database

import (
	"go-project/internal/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	errMigrate := db.AutoMigrate(&models.User{})
	if errMigrate != nil {
		panic(errMigrate)
	}
}
