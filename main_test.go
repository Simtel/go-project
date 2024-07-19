package main

import (
	"go-project/models"
	"testing"
)

func TestUser(t *testing.T) {
	u := models.User{Name: "Simtel", Email: "email@email.com"}

	if u.Name != "Simtel" {
		t.Errorf("Expected name to be 'Simtel', but got '%s'", u.Name)
	}

	if u.Email != "email@email.com" {
		t.Errorf("Expected email to be 'email@email.com', but got '%s'", u.Email)
	}

	if u.GetName() != "Simtel" {
		t.Errorf("Expected name to be 'Simtel', but got '%s'", u.GetName())
	}
}
