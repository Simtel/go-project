package api

import (
	"bytes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/internal/common"
	"html/template"
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

	a.r.Get("/template", func(w http.ResponseWriter, r *http.Request) {

		data := struct {
			Title string
			Name  string
		}{
			Title: "Пример шаблона",
			Name:  "Мир",
		}

		tmpl, err := template.ParseFiles("internal/resource/templates/layout.html")
		if err != nil {
			common.SendErrorResponse(w, err.Error())
		}
		var buf bytes.Buffer

		errExecute := tmpl.Execute(&buf, data)
		if errExecute != nil {
			common.SendErrorResponse(w, err.Error())
		}

		result := buf.String()

		common.SendSuccessJsonResponse(w, result)
	})
}

type Home struct {
	Payload string `json:"payload"`
	Success bool   `json:"success"`
}
