package domainsrepo

import (
	"go-project/internal/models"
	"go-project/internal/services/armisimtel"
)

type HttpRepositoryInterface interface {
	GetByName(name string) (*models.Domain, error)
	GetAll(c chan []*models.Domain) ([]*models.Domain, error)
	GetById(domainId int) (*models.Domain, error)
	New(payload *armisimtel.DomainPayload) (*models.Domain, error)
}
