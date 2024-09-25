package domainsrepo

import "go-project/internal/models"

type MysqlRepositoryInterface interface {
	Create(domain *models.Domain)
}
