package models

type Domain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ExpireAt string `json:"expire_at"`
}

func GetId(domain *Domain) int {
	return domain.ID
}
