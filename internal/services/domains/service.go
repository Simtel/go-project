package domains

import "go-project/internal/adapter/httprepo/domainsrepo"

type Service struct {
	repo *domainsrepo.Repository
}

func NewDomainsService(repo *domainsrepo.Repository) *Service {
	return &Service{repo}
}

func (s *Service) ShowDomains() {

}
