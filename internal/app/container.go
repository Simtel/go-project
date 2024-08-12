package app

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/adapter/httprepo/domainsrepo"
	"go-project/internal/adapter/mysqlrepo/usersrepo"
	"go-project/internal/handler/http/api"
	"go-project/internal/services/armisimtel"
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

func (c *Container) GetDomainsRepo() *domainsrepo.Repository {
	return domainsrepo.NewRepository(c.GetArmiSimtelRequest())
}

func (c *Container) GetArmiSimtelRequest() armisimtel.RequestInterface {
	return armisimtel.NewRequest(c.GetHttpClient())
}

func (c *Container) GetDomainsApi() *api.DomainsApi {
	return api.NewDomainsApi(c.GetRouter(), c.GetDomainsRepo())
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
