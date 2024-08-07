package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-gormigrate/gormigrate/v2"
	"go-project/internal/app"
	"go-project/internal/common"
	"go-project/internal/database"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	common.InitEnv()
	common.InitFileStorage()

	db := database.NewDbMysql()
	database.MigrateDB(db)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{{
		ID: "202208301400",
		Migrate: func(tx *gorm.DB) error {
			type contact struct {
				ID   uint `gorm:"primaryKey;uniqueIndex"`
				Name string
			}
			return tx.Migrator().CreateTable(&contact{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("contacts")
		},
	}, {
		ID: "202208301415",
		Migrate: func(tx *gorm.DB) error {

			type contact struct {
				Parent int
			}
			return tx.Migrator().AddColumn(&contact{}, "Parent")
		},
		Rollback: func(tx *gorm.DB) error {
			type user struct {
				Age int
			}
			return db.Migrator().DropColumn(&user{}, "Age")
		},
	},
	})

	if errMigrate := m.Migrate(); errMigrate != nil {
		log.Fatalf("Migration failed: %v", errMigrate)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	a := app.NewContainer(&http.Client{}, r)

	a.AddHandler(a.GetDomainsApi())
	a.AddHandler(a.GetMainApi())
	a.AddHandler(a.GetUsersApi(db))

	errServe := http.ListenAndServe(":3000", r)
	if errServe != nil {
		panic(errServe)
	}
}
