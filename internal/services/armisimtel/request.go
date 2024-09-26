package armisimtel

import (
	"errors"
	"net/http"
	"os"
)

var BASE_URL string = "https://armisimtel.ru/api/v1"

var Token string

type Request struct {
	client *http.Client
}

//go:generate mockgen -source=request.go -destination=../../../mock/request.go -package=mock
type RequestInterface interface {
	Request(method string, url string, data []byte) (*http.Response, error)
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
