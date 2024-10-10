package app

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/adapter/httprepo/domainsrepo"
	mysqldomainsrepo "go-project/internal/adapter/mysqlrepo/domainsrepo"
	"go-project/internal/adapter/mysqlrepo/usersrepo"
	"go-project/internal/handler/http/api"
	"go-project/internal/handler/http/api/domains"
	"go-project/internal/services/armisimtel"
	"go-project/internal/services/domains/storage"
	"gorm.io/gorm"
	"net/http"
)

type Container struct {
	http   *http.Client
	router *chi.Mux
	db     *gorm.DB
}

func NewContainer(http *http.Client, router *chi.Mux, db *gorm.DB) *Container {
	return &Container{
		http,
		router,
		db,
	}
}

func (c *Container) GetDomainsRepo() domainsrepo.HttpRepositoryInterface {
	return domainsrepo.NewRepository(c.GetArmiSimtelRequest())
}

func (c *Container) GetArmiSimtelRequest() armisimtel.RequestInterface {
	return armisimtel.NewRequest(c.GetHttpClient())
}

func (c *Container) GetDomainsApi() *domains.DomainsApi {
	return domains.NewDomainsApi(c.GetRouter(), c.GetDomainsRepo(), c.GetMysqlDomainsRepo(), c.GetDomainStorage())
}

func (c *Container) GetHttpClient() *http.Client {
	return c.http
}

func (c *Container) GetRouter() *chi.Mux {
	return c.router
}

func (c *Container) GetMainApi() *api.MainApi {
	return api.NewMainApi(c.GetRouter())
}

func (c *Container) GetUsersApi() *api.UsersApi {
	return api.NewUsersApi(c.GetRouter(), usersrepo.NewUsersRepo(c.GetDB()))
}

func (c *Container) GetDB() *gorm.DB {
	return c.db
}

func (c *Container) AddHandler(h api.Handler) {
	h.AddRoutes()
}

func (c *Container) GetMysqlDomainsRepo() mysqldomainsrepo.MysqlRepositoryInterface {
	return mysqldomainsrepo.NewDomainsRepo(c.GetDB())
}

func (c *Container) GetDomainStorage() storage.DomainStorageInterface {
	return storage.NewStorageDomain()
}
