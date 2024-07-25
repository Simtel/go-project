package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/common"
	"go-project/domains"
	"net/http"
)

func main() {
	common.InitEnv()
	common.InitFileStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	domains.Routes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		common.SendSuccessJsonResponse(w, "Hello")
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		common.SendErrorResponse(w, "Something went wrong")
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
