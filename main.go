package main

import (
	"github.com/go-chi/chi/v5"
	"go-project/common"
	"go-project/domains"
	"net/http"
)

func main() {
	r := chi.NewRouter()

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
