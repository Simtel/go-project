package domainsrepo

import (
	"gorm.io/gorm"
)

type DomainsRepo struct {
	db *gorm.DB
}

func NewDomainsRepo(db *gorm.DB) *DomainsRepo {
	return &DomainsRepo{db: db}
}
