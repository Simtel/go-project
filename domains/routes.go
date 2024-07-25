package domains

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/common"
	"net/http"
	"os"
	"strconv"
)

func Routes(r *chi.Mux) {

	r.Get("/domains", func(w http.ResponseWriter, r *http.Request) {
		domainsList, err := ShowDomains()
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}
		go func() {
			err := SaveDomains(domainsList)
			if err != nil {
				_ = fmt.Errorf("eror save fomain in file: %s", err)
			}
		}()

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

	r.Get("/domains/download", func(w http.ResponseWriter, r *http.Request) {
		file, errOpen := os.Open("var/domains.csv")
		if errOpen != nil {
			common.SendErrorResponse(w, errOpen.Error())
			return
		}
		defer file.Close()

		common.SendFile(w, r, file)

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
