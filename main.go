package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/app"
	"go-project/common"
	"net/http"
)

func main() {
	common.InitEnv()
	common.InitFileStorage()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	a := app.NewContainer(&http.Client{}, r)

	a.GetDomainsApi().AddRoutes()
	a.GetMainApi().AddRoutes()

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
