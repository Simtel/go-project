package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/internal/app"
	"go-project/internal/common"
	"go-project/internal/database"
	"net/http"
)

func main() {
	common.InitEnv()
	common.InitFileStorage()

	db := database.NewDbMysql()
	database.MigrateDB(db)

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
