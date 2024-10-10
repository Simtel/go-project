package domains

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/adapter/httprepo/domainsrepo"
	mysqldomainsrepo "go-project/internal/adapter/mysqlrepo/domainsrepo"
	"go-project/internal/services/domains/storage"
)

type DomainsApi struct {
	r         *chi.Mux
	httpRepo  domainsrepo.HttpRepositoryInterface
	mysqlRepo mysqldomainsrepo.MysqlRepositoryInterface
	storage   storage.DomainStorageInterface
}

func NewDomainsApi(
	r *chi.Mux,
	httpRepo domainsrepo.HttpRepositoryInterface,
	mysqlRepo mysqldomainsrepo.MysqlRepositoryInterface,
	storage storage.DomainStorageInterface,
) *DomainsApi {
	return &DomainsApi{r: r, httpRepo: httpRepo, mysqlRepo: mysqlRepo, storage: storage}
}

func (a *DomainsApi) AddRoutes() {
	a.r.Get("/domains", a.GetDomains)
	a.r.Get("/domains/{id}", a.GetDomainById)
	a.r.Post("/domains", a.CreateDomain)
	a.r.Get("/domains/download", a.Download)
	a.r.Get("/domains/local", a.LocalDomains)
}
