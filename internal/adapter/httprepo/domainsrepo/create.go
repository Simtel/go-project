package domainsrepo

import (
	"encoding/json"
	"errors"
	"go-project/internal/models"
	"go-project/internal/services/armisimtel"
)

func (r *Repository) Create(payload *armisimtel.DomainPayload) (*models.Domain, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("Ошибка преобразования JSON:" + err.Error())
	}
	resp, errorResp := r.request.Request("POST", "/domains", jsonData)
	if errorResp != nil {
		return nil, errorResp
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	return &models.Domain{Name: payload.Name}, nil
}
