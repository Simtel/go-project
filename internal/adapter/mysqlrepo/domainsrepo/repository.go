package domainsrepo

import (
	"gorm.io/gorm"
)

type DomainsRepo struct {
	db *gorm.DB
}

func NewDomainsRepo(db *gorm.DB) MysqlRepositoryInterface {
	return &DomainsRepo{db: db}
}
