package api

import (
	"bytes"
	"encoding/json"
	"go-project/internal/services/armisimtel"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"go-project/internal/models"
)

type MockRepo struct{}

func (m *MockRepo) GetAll(ch chan []*models.Domain) ([]*models.Domain, error) {
	domains := []*models.Domain{
		{ID: 1, Name: "example.com", ExpireAt: "2023-12-31"},
		{ID: 2, Name: "test.com", ExpireAt: "2024-01-31"},
	}
	ch <- domains
	return domains, nil
}

func (m *MockRepo) GetById(id int) (*models.Domain, error) {
	return &models.Domain{ID: id, Name: "example.com", ExpireAt: "2023-12-31"}, nil
}
func (m *MockRepo) GetByName(name string) (*models.Domain, error) {
	return &models.Domain{ID: 1, Name: name, ExpireAt: "2023-12-31"}, nil
}

func (m *MockRepo) New(payload *armisimtel.DomainPayload) (*models.Domain, error) {
	return &models.Domain{ID: 3, Name: "new.com", ExpireAt: "2025-01-01"}, nil
}

type MockMySQLRepo struct{}

func (m *MockMySQLRepo) Create(domain *models.Domain) {

}

func TestDomainsApi_GetAllDomains(t *testing.T) {
	r := chi.NewRouter()

	mockMySQLRepo := &MockMySQLRepo{}
	api := NewDomainsApi(r, &MockRepo{}, mockMySQLRepo)
	api.AddRoutes()

	req, err := http.NewRequest("GET", "/domains", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Status  bool             `json:"status"`
		Payload []*models.Domain `json:"payload"`
		Message string           `json:"message"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Status {
		t.Errorf("expected success to be true")
	}
	if len(response.Payload) != 2 {
		t.Errorf("expected 2 domains, got %d", len(response.Payload))
	}
}

func TestDomainsApi_GetDomainById(t *testing.T) {
	r := chi.NewRouter()
	mockRepo := &MockRepo{}
	mockMySQLRepo := &MockMySQLRepo{}
	api := NewDomainsApi(r, mockRepo, mockMySQLRepo)
	api.AddRoutes()

	req, err := http.NewRequest("GET", "/domains/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Status  bool           `json:"status"`
		Payload *models.Domain `json:"payload"`
		Message string         `json:"message"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Status {
		t.Errorf("expected success to be true")
	}
	if response.Payload.ID != 1 {
		t.Errorf("expected domain ID 1, got %d", response.Payload.ID)
	}
}

func TestDomainsApi_CreateDomain(t *testing.T) {
	r := chi.NewRouter()
	mockRepo := &MockRepo{}
	mockMySQLRepo := &MockMySQLRepo{}
	api := NewDomainsApi(r, mockRepo, mockMySQLRepo)
	api.AddRoutes()

	payload := []byte(`{"name": "new.com"}`)
	req, err := http.NewRequest("POST", "/domains", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Status  bool           `json:"status"`
		Payload *models.Domain `json:"payload"`
		Message string         `json:"message"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Status {
		t.Errorf("expected success to be true")
	}
	if response.Payload.ID != 3 {
		t.Errorf("expected new domain ID 3, got %d", response.Payload.ID)
	}
}

func TestDomainsApi_DownloadDomains(t *testing.T) {
	r := chi.NewRouter()
	mockRepo := &MockRepo{}
	mockMySQLRepo := &MockMySQLRepo{}
	api := NewDomainsApi(r, mockRepo, mockMySQLRepo)
	api.AddRoutes()

	req, err := http.NewRequest("GET", "/domains/download", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "text/csv; charset=utf-8" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/csv; charset=utf-8")
	}
}
