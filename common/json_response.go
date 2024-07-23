package common

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Payload any    `json:"payload"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func NewJsonResponse(payload any, message string, status bool) JsonResponse {
	return JsonResponse{payload, message, status}
}

func SendSuccessJsonResponse(w http.ResponseWriter, payload any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(NewJsonResponse(payload, "", true))
	if err != nil {
		panic(err)
	}
}

func SendForbiddenResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(NewJsonResponse("", message, true))
	if err != nil {
		panic(err)
	}
}

func SendErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(NewJsonResponse("", message, true))
	if err != nil {
		panic(err)
	}
}
