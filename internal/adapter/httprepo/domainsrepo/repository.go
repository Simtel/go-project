package domainsrepo

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-project/internal/models"
	"go-project/internal/services/armisimtel"
	"io"
	"strconv"
)

type Repository struct {
	request armisimtel.RequestInterface
}

func NewRepository(request armisimtel.RequestInterface) *Repository {
	return &Repository{
		request: request,
	}
}

func (r *Repository) GetByName(name string) (*models.Domain, error) {
	return nil, errors.New("not implemented")
}

func (r *Repository) GetAll() ([]*models.Domain, error) {

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
	return parsed.Data, nil
}

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

func (r *Repository) New(payload *armisimtel.DomainPayload) (*models.Domain, error) {
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
