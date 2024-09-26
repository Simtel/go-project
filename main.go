package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-project/internal/app"
	"go-project/internal/config"
	"go-project/internal/database"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// инициализация среды и хранилища файлов
func initialize() {
	config.InitEnv()
	config.InitFileStorage()
}

// настройка базы данных и выполнение миграций
func setupDatabase() *gorm.DB {
	db := database.NewDbMysql()
	database.MigrateDB(db)

	m := gormigrate.New(db, gormigrate.DefaultOptions, database.GetMigrations())

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}

// создание и настройка маршрутизатора
func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	return r
}

// основной запуск приложения
func main() {
	initialize()

	db := setupDatabase()

	r := setupRouter()

	a := app.NewContainer(&http.Client{}, r, db)

	// добавление обработчиков маршрутов
	a.AddHandler(a.GetDomainsApi())
	a.AddHandler(a.GetMainApi())
	a.AddHandler(a.GetUsersApi())

	log.Println("Server is starting on port 3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
