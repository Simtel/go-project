package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDbMysql() *gorm.DB {
	mysqlDsn, exists := os.LookupEnv("MYSQL_DSN")
	if !exists {
		panic("MYSQL environment variable not set")
	}
	db, errConn := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if errConn != nil {
		panic(errConn)
	}

	return db
}
