package domainsrepo

import (
	"go-project/common"
	"go-project/internal/services/armisimtel"
	"net/http"
	"os"
	"testing"
)

func setup() {
	common.InitEnv()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestRepo(t *testing.T) {

	request := armisimtel.NewRequest(&http.Client{})
	repo := NewRepository(request)

	domains, err := repo.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(domains) == 0 {
		t.Error("domains is empty")
	}
}
