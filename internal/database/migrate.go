package database

import (
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	errMigrate := db.AutoMigrate(&db.User{})
	if errMigrate != nil {
		panic(errMigrate)
	}
}
