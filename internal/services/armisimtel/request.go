package armisimtel

import (
	"bytes"
	"errors"
	"net/http"
	"os"
)

var BASE_URL string = "https://armisimtel.ru/api/v1"

var Token string

type Request struct {
	client *http.Client
}

func NewRequest(client *http.Client) *Request {
	return &Request{
		client: client,
	}
}

func (request *Request) Init() (string, error) {

	token, exists := os.LookupEnv("ARMISIMTEL_TOKEN")
	if !exists {
		return "", errors.New("you need to set ARMISIMTEL_TOKEN environment variable")
	}
	return token, nil
}

func (request *Request) Request(method string, url string, data []byte) (*http.Response, error) {
	token, err := request.Init()
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token

	req, errorReq := http.NewRequest(method, BASE_URL+url, bytes.NewBuffer(data))

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
