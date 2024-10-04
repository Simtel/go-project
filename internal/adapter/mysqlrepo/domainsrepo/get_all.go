package domainsrepo

import (
	"errors"
	"go-project/internal/models/db"
)

func (r *DomainsRepo) GetAll() ([]*db.Domain, error) {
	var domains []*db.Domain
	result := r.db.Find(&domains)
	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}
	return domains, nil
}
