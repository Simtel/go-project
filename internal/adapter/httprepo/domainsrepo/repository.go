package domainsrepo

import (
	"go-project/internal/services/armisimtel"
)

type Repository struct {
	request armisimtel.RequestInterface
}

func NewRepository(request armisimtel.RequestInterface) *Repository {
	return &Repository{
		request: request,
	}
}
