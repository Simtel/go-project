package domainsrepo

import (
	"bytes"
	"errors"
	gomock2 "github.com/golang/mock/gomock"
	"go-project/internal/services/armisimtel"
	"go-project/mock"
	"io"
	"net/http"
	"testing"
)

func TestRepo(t *testing.T) {
	gomock := &gomock2.Controller{T: t}

	request := MockRequest(gomock, SimulateResponse(), errors.New("some error"))
	repo := NewRepository(request)

	domains, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(domains) == 0 {
		t.Error("domains is empty")
	}
}

func MockRequest(controller *gomock2.Controller, result *http.Response, err error) armisimtel.RequestInterface {
	mockRequest := mock.NewMockRequestInterface(controller)
	mockRequest.EXPECT().Request("GET", "/domains", nil).Return(result, err)

	return mockRequest
}

func SimulateResponse() *http.Response {
	responseBody := io.NopCloser(bytes.NewReader([]byte(`{"value":"fixed"}`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responseBody,
	}
}
