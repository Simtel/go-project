package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go-project/internal/adapter/httprepo/domainsrepo"
	mysqldomainsrepo "go-project/internal/adapter/mysqlrepo/domainsrepo"
	"go-project/internal/common"
	"go-project/internal/models"
	"go-project/internal/services/armisimtel"
	domains2 "go-project/internal/services/domains"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
)

type DomainsApi struct {
	r         *chi.Mux
	httpRepo  domainsrepo.HttpRepositoryInterface
	mysqlRepo mysqldomainsrepo.MysqlRepositoryInterface
}

func NewDomainsApi(r *chi.Mux, httpRepo domainsrepo.HttpRepositoryInterface, mysqlRepo mysqldomainsrepo.MysqlRepositoryInterface) *DomainsApi {
	return &DomainsApi{r: r, httpRepo: httpRepo, mysqlRepo: mysqlRepo}
}

func (a *DomainsApi) AddRoutes() {

	a.r.Get("/domains", func(w http.ResponseWriter, r *http.Request) {
		domainsChannel := make(chan []*models.Domain)

		go func(c chan []*models.Domain) {
			domainsList := <-c
			err := domains2.SaveDomains(domainsList, "var/api.csv")
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
	})

	a.r.Get("/domains/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	a.r.Post("/domains", func(w http.ResponseWriter, r *http.Request) {
		domain := &armisimtel.DomainPayload{}
		if err := render.Bind(r, domain); err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}

		createDomain, err := a.httpRepo.New(domain)
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}

		common.SendSuccessJsonResponse(w, createDomain)

	})

	a.r.Get("/domains/download", func(w http.ResponseWriter, r *http.Request) {
		_, b, _, _ := runtime.Caller(0)
		d1 := path.Join(path.Dir(b))
		file, errOpen := os.Open(d1 + "/../../../../var/api.csv")
		if errOpen != nil {
			common.SendErrorResponse(w, errOpen.Error())
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}(file)

		common.SendFile(w, r, file)

	})
}
