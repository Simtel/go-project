package api

import (
	"bytes"
	"github.com/domodwyer/mailyak"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/internal/common"
	"html/template"
	"net/http"
	"net/smtp"
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

	a.r.Get("/slice", func(w http.ResponseWriter, r *http.Request) {
		var slice []string

		slice = append(slice, "test")
		slice = append(slice, "example")
		slice = append(slice, "1")
		common.SendSuccessJsonResponse(w, slice)
	})

	a.r.Get("/sendmail", func(w http.ResponseWriter, r *http.Request) {
		mail := mailyak.New("localhost:1025", smtp.PlainAuth("", "", "", "localhost"))
		mail.To("dom@itsallbroken.com")
		mail.From("jsmith@example.com")
		mail.FromName("Bananas for Friends")
		mail.Subject("Business proposition")
		// Or set the body using a string setter
		mail.Plain().Set("Get a real email client")

		// And you're done!
		if err := mail.Send(); err != nil {
			panic(err)
		}
	})

	a.r.Get("/ip", func(w http.ResponseWriter, r *http.Request) {
		type Result struct {
			Ip string `json:"ip"`
		}
		var result *Result
		forwarded := r.Header.Get("X-FORWARDED-FOR")
		if forwarded != "" {
			result = &Result{Ip: forwarded}

		} else {
			result = &Result{Ip: r.RemoteAddr}
		}

		common.SendSuccessJsonResponse(w, result)
	})
}

type Home struct {
	Payload string `json:"payload"`
	Success bool   `json:"success"`
}
