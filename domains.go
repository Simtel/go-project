package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

var Token string

type Domain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ExpireAt string `json:"expire_at"`
}

func ShowDomains() ([]Domain, error) {
	token, err := Init()
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token

	req, errorReq := http.NewRequest("GET", "https://armisimtel.ru/api/v1/domains", bytes.NewBuffer(nil))

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

	var parsed struct {
		Data []Domain `json:"data"`
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	errJson := json.Unmarshal([]byte(body), &parsed)
	if errJson != nil {

		return nil, errors.New("Ошибка декодирования JSON:" + err.Error())
	}

	return parsed.Data, nil
}

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
