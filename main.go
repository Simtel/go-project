package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		SendSuccessJsonResponse(w, "Hello")
		return
	})

	r.Get("/domains", func(w http.ResponseWriter, r *http.Request) {
		domains, err := ShowDomains()
		if err != nil {
			SendErrorResponse(w, err.Error())
			return
		}
		SendSuccessJsonResponse(w, domains)
		return
	})

	r.Get("/domains/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		domainId, errConvert := strconv.Atoi(chi.URLParam(r, "id"))
		if errConvert != nil {
			SendErrorResponse(w, errConvert.Error())
			return
		}
		domains, err := ShowDomainById(domainId)
		if err != nil {
			SendErrorResponse(w, err.Error())
			return
		}
		SendSuccessJsonResponse(w, domains)
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		SendErrorResponse(w, "Something went wrong")
		return
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
