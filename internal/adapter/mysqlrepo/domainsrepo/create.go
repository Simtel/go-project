package domainsrepo

import (
	"fmt"
	"go-project/internal/models"
	"go-project/internal/models/db"
	"time"
)

func (r *DomainsRepo) Create(domain *models.Domain) {
	layout := time.DateTime
	expireTime, err := time.Parse(layout, domain.ExpireAt)
	if err != nil {
		fmt.Println("Ошибка парсинга времени истекания:", err)
		return
	}

	model := db.Domain{Domain: domain.Name, User: 1, Expired: expireTime, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	r.db.Create(&model)
}
