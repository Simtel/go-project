package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(NewJsonResponse("Hello", "", true))
		return
	})

	r.Get("/domains", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		domains, err := ShowDomains()
		if err != nil {
			SendErrorResponse(w, err.Error())
		}
		SendSuccessJsonResponse(w, domains)
	})

	http.ListenAndServe(":3000", r)
}
