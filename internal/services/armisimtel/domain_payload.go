package armisimtel

import (
	"errors"
	"net/http"
)

type DomainPayload struct {
	Name string `json:"name"`
}

func (d *DomainPayload) Bind(r *http.Request) error {

	if d.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
