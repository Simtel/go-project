package domainsrepo

import (
	"go-project/internal/models"
	"go-project/internal/models/db"
)

type MysqlRepositoryInterface interface {
	Create(domain *models.Domain)
	GetAll() ([]*db.Domain, error)
}
