package main

import "testing"

func TestNewJsonResponse(t *testing.T) {
	jsonResponse := NewJsonResponse("test", "message string", true)

	if jsonResponse.Payload != "test" {
		t.Error("Expected test in payload, got", jsonResponse.Payload)
	}

	if jsonResponse.Message != "message string" {
		t.Error("Expected message string in message, got", jsonResponse.Message)
	}

	if jsonResponse.Status != true {
		t.Error("Expected true in status, got", jsonResponse.Status)
	}
}
