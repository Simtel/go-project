package database

import (
	db2 "go-project/internal/models/db"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	errMigrate := db.AutoMigrate(&db2.User{})
	if errMigrate != nil {
		panic(errMigrate)
	}
}
