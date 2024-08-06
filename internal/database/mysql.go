package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDbMysql() *gorm.DB {
	mysql_dsn, exists := os.LookupEnv("MYSQL_DSN")
	if !exists {
		panic("MYSQL environment variable not set")
	}
	db, errConn := gorm.Open(mysql.Open(mysql_dsn), &gorm.Config{})
	if errConn != nil {
		panic(errConn)
	}

	return db
}
