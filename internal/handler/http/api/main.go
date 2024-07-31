package api

import (
	"github.com/go-chi/chi/v5"
	"go-project/common"
	"net/http"
)

type MainApi struct {
	r *chi.Mux
}

func NewMainApi(r *chi.Mux) *MainApi {
	return &MainApi{r: r}
}

func (a *MainApi) AddRoutes() {

	a.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		common.SendSuccessJsonResponse(w, "Hello, World")
	})

	a.r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		common.SendErrorResponse(w, "Something went wrong")
	})
}
