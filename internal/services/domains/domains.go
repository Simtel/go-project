package domains

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-project/models"
	"io"
	"net/http"
	"os"
	"strconv"
)

var BASE_URL string = "https://armisimtel.ru/api/v1"

var Token string

func Init() (string, error) {

	token, exists := os.LookupEnv("ARMISIMTEL_TOKEN")
	if !exists {
		return "", errors.New("you need to set ARMISIMTEL_TOKEN environment variable")
	}
	return token, nil
}

func ShowDomains() ([]*models.Domain, error) {

	resp, errorResp := request("GET", BASE_URL+"/domains", nil)

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

func ShowDomainById(domainId int) (*models.Domain, error) {
	resp, errorResp := request("GET", BASE_URL+"/domains/"+strconv.Itoa(domainId), nil)

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

func CreateDomain(payload *DomainPayload) (*models.Domain, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("Ошибка преобразования JSON:" + err.Error())
	}
	resp, errorResp := request("POST", BASE_URL+"/domains", jsonData)
	if errorResp != nil {
		return nil, errorResp
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	return &models.Domain{Name: payload.Name}, nil
}

func request(method string, url string, data []byte) (*http.Response, error) {
	token, err := Init()
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token

	req, errorReq := http.NewRequest(method, url, bytes.NewBuffer(data))

	if errorReq != nil {
		return nil, errorReq
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	resp, errorResp := client.Do(req)

	if errorResp != nil {
		return nil, errorResp
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return resp, errorResp
}

type DomainPayload struct {
	Name string `json:"name"`
}

func (d *DomainPayload) Bind(r *http.Request) error {

	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
