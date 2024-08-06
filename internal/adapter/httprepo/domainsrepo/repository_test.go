package domainsrepo

import (
	"go-project/internal/common"
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

	domains, err := repo.GetByName("test.com")

	if domains != nil {
		t.Error("domains should be nil")
	}

	if err.Error() != "not implemented" {
		t.Error(err)
	}

}
