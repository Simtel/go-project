package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/app"
	"go-project/common"
	"go-project/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	common.InitEnv()
	mysql_dsn, exists := os.LookupEnv("MYSQL_DSN")
	if !exists {
		panic("MYSQL environment variable not set")
	}
	db, errConn := gorm.Open(mysql.Open(mysql_dsn), &gorm.Config{})
	if errConn != nil {
		panic(errConn)
	}

	errMigrate := db.AutoMigrate(&models.User{})
	if errMigrate != nil {
		panic(errMigrate)
	}

	common.InitFileStorage()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	a := app.NewContainer(&http.Client{}, r)

	a.GetDomainsApi().AddRoutes()
	a.GetMainApi().AddRoutes()

	errServe := http.ListenAndServe(":3000", r)
	if errServe != nil {
		panic(errServe)
	}
}
