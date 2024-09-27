package app

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-project/internal/config"
	"go-project/internal/database"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	c *Container
}

var app *App

func NewApp() (*App, error) {
	config.InitEnv()
	config.InitFileStorage()

	db := setupDatabase()
	r := setupRouter()

	a := NewContainer(&http.Client{}, r, db)

	return &App{c: a}, nil
}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	return r
}

func setupDatabase() *gorm.DB {
	db := database.NewDbMysql()
	database.MigrateDB(db)

	m := gormigrate.New(db, gormigrate.DefaultOptions, database.GetMigrations())

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}

func SetApplication(a *App) {
	app = a
}

func GetApplication() (*App, error) {
	if app == nil {
		return nil, errors.New("global app is not initialized")
	}

	return app, nil
}

func (a *App) Container() *Container {
	return a.c
}
