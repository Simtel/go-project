package domainsrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-project/internal/models"
	"io"
)

func (r *Repository) GetAll(c chan []*models.Domain) ([]*models.Domain, error) {

	resp, errorResp := r.request.Request("GET", "/domains", nil)

	if errorResp != nil {
		return nil, errorResp
	}

	var parsed struct {
		Data []*models.Domain `json:"data"`
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	errJson := json.Unmarshal([]byte(body), &parsed)
	if errJson != nil {

		return nil, errors.New("Ошибка декодирования JSON:" + err.Error())
	}
	fmt.Println("Получен общий список доменов")
	c <- parsed.Data
	return parsed.Data, nil
}
