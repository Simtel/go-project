package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func ShowDomains() {
	Init()

	fmt.Println("Token:", Token)
	bearer := "Bearer " + Token

	req, errorReq := http.NewRequest("GET", "https://armisimtel.ru/api/v1/domains", bytes.NewBuffer(nil))

	if errorReq != nil {
		fmt.Println(errorReq)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	resp, errorResp := client.Do(req)

	if errorResp != nil {
		fmt.Println(errorResp)
	}

	fmt.Println(resp.Status)

	var parsed struct {
		Data []Domain `json:"data"`
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	errJson := json.Unmarshal([]byte(body), &parsed)
	if errJson != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	for _, domain := range parsed.Data {
		fmt.Println(domain.Name)
	}
}

func Init() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		fmt.Println("Error load env file")
	}

	token, exists := os.LookupEnv("ARMISIMTEL_TOKEN")
	if !exists {
		fmt.Println("You need to set ARMISIMTEL_TOKEN environment variable")
	}
	fmt.Println("Token:", token)
	Token = token
}
