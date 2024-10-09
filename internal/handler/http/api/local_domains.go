package api

import (
	"go-project/internal/common"
	"net/http"
)

func (a *DomainsApi) LocalDomains(w http.ResponseWriter, r *http.Request) {
	domains, err := a.mysqlRepo.GetAll()
	if err != nil {
		common.SendErrorResponse(w, err.Error())
	}
	common.SendSuccessJsonResponse(w, domains)
}
