package models

type Domain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ExpireAt string `json:"expire_at"`
}

func (domain *Domain) GetId() int {
	return domain.ID
}
