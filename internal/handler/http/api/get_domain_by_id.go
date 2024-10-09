package api

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/common"
	"net/http"
	"strconv"
)

func (a *DomainsApi) GetDomainById(w http.ResponseWriter, r *http.Request) {

	domainId, errConvert := strconv.Atoi(chi.URLParam(r, "id"))
	if errConvert != nil {
		common.SendErrorResponse(w, errConvert.Error())
		return
	}
	domain, err := a.httpRepo.GetById(domainId)
	if err != nil {
		common.SendErrorResponse(w, err.Error())
		return
	}
	common.SendSuccessJsonResponse(w, domain)
}
