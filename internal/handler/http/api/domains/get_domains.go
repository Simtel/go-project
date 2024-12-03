package domains

import (
	"fmt"
	"go-project/internal/common"
	"go-project/internal/models"
	"net/http"
)

func (a *Api) GetDomains(w http.ResponseWriter, r *http.Request) {
	domainsChannel := make(chan []*models.Domain)

	go func(c chan []*models.Domain) {
		domainsList := <-c
		err := a.storage.Save(domainsList, "var/api.csv")
		if err != nil {
			_ = fmt.Errorf("eror save fomain in file: %s", err)
		}
	}(domainsChannel)

	domainsList, err := a.httpRepo.GetAll(domainsChannel)
	if err != nil {
		common.SendErrorResponse(w, err.Error())
		return
	}

	for _, domain := range domainsList {
		a.mysqlRepo.Create(domain)
	}

	common.SendSuccessJsonResponse(w, domainsList)
}
