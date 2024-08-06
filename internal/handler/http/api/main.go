package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/internal/common"
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
		render.JSON(w, r, Home{Payload: "Hello", Success: true})
	})

	a.r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		common.SendErrorResponse(w, "Something went wrong")
	})
}

type Home struct {
	Payload string `json:"payload"`
	Success bool   `json:"success"`
}
