package main

import (
	"go-project/models"
	"testing"
)

func TestUser(t *testing.T) {
	u := models.User{name: "Simtel"}

	if u.name != "Simtel" {
		t.Errorf("Expected name to be 'Simtel', but got '%s'", u.name)
	}

	if u.email != "simtel@gmail.com" {
		t.Errorf("Expected email to be 'simtel@gmail.com', but got '%s'", u.email)
	}
}
