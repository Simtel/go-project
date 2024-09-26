package domainsrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-project/internal/models"
	"io"
	"strconv"
)

func (r *Repository) GetById(domainId int) (*models.Domain, error) {
	resp, errorResp := r.request.Request("GET", "/domains/"+strconv.Itoa(domainId), nil)

	if errorResp != nil {
		return nil, errorResp
	}

	var parsed struct {
		Data models.Domain `json:"data"`
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	errJson := json.Unmarshal([]byte(body), &parsed)
	if errJson != nil {
		return nil, errors.New("Ошибка декодирования JSON:" + err.Error())
	}
	fmt.Println("Получена информация по домену с ID: " + strconv.Itoa(domainId))
	return &parsed.Data, nil
}
