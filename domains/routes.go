package domains

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/common"
	"net/http"
	"strconv"
)

func Routes(r *chi.Mux) {

	r.Get("/domains", func(w http.ResponseWriter, r *http.Request) {
		domainsList, err := ShowDomains()
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}
		common.SendSuccessJsonResponse(w, domainsList)
	})

	r.Get("/domains/{id}", func(w http.ResponseWriter, r *http.Request) {
		domainId, errConvert := strconv.Atoi(chi.URLParam(r, "id"))
		if errConvert != nil {
			common.SendErrorResponse(w, errConvert.Error())
			return
		}
		domain, err := ShowDomainById(domainId)
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}
		common.SendSuccessJsonResponse(w, domain)
	})

	r.Post("/domains", func(w http.ResponseWriter, r *http.Request) {
		domain := &DomainPayload{}
		if err := render.Bind(r, domain); err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}

		createDomain, err := CreateDomain(domain)
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}

		common.SendSuccessJsonResponse(w, createDomain)

	})
}

type DomainPayload struct {
	Name string `json:"name"`
}

func (d *DomainPayload) Bind(r *http.Request) error {

	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
