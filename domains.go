package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go-project/models"
	"io"
	"net/http"
	"os"
	"strconv"
)

var BASE_URL string = "https://armisimtel.ru/api/v1"

var Token string

func Init() (string, error) {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		return "", errors.New("error load env file")
	}

	token, exists := os.LookupEnv("ARMISIMTEL_TOKEN")
	if !exists {
		return "", errors.New("you need to set ARMISIMTEL_TOKEN environment variable")
	}
	return token, nil
}

func ShowDomains() ([]models.Domain, error) {

	resp, errorResp := request(BASE_URL + "/domains")

	if errorResp != nil {
		return nil, errorResp
	}

	var parsed struct {
		Data []models.Domain `json:"data"`
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
	resp, errorResp := request(BASE_URL + "/domains/" + strconv.Itoa(domainId))

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

func request(url string) (*http.Response, error) {
	token, err := Init()
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token

	req, errorReq := http.NewRequest("GET", url, bytes.NewBuffer(nil))

	if errorReq != nil {
		return nil, errorReq
	}

	req.Header.Add("Authorization", bearer)
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
