package domains

import (
	"github.com/go-chi/render"
	"go-project/internal/common"
	"go-project/internal/services/armisimtel"
	"net/http"
)

func (a *Api) CreateDomain(w http.ResponseWriter, r *http.Request) {
	domain := &armisimtel.DomainPayload{}
	if err := render.Bind(r, domain); err != nil {
		common.SendErrorResponse(w, err.Error())
		return
	}

	createDomain, err := a.httpRepo.Create(domain)
	if err != nil {
		common.SendErrorResponse(w, err.Error())
		return
	}

	common.SendSuccessJsonResponse(w, createDomain)
}
