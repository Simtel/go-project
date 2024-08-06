package api

import (
	"github.com/go-chi/chi/v5"
	"go-project/internal/adapter/mysqlrepo/usersrepo"
	"go-project/internal/common"
	"net/http"
)

type UsersApi struct {
	r    *chi.Mux
	repo *usersrepo.UsersRepo
}

func NewUsersApi(r *chi.Mux, repo *usersrepo.UsersRepo) *UsersApi {
	return &UsersApi{r: r, repo: repo}
}

func (a UsersApi) AddRoutes() {
	a.r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.repo.GetAll()
		if err != nil {
			common.SendErrorResponse(w, err.Error())
			return
		}
		common.SendSuccessJsonResponse(w, users)
	})
}
