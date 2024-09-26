package storage

import (
	"go-project/internal/models"
	"os"
)

type DomainStorage struct {
}

func NewStorageDomain() DomainStorageInterface {
	return &DomainStorage{}
}

type DomainStorageInterface interface {
	Save(domains []*models.Domain, filePath string) error
	Get(filepath string) (*os.File, error)
}
