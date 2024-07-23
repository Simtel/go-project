package common

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Payload interface{} `json:"payload"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

func NewJsonResponse(payload interface{}, message string, status bool) JsonResponse {
	return JsonResponse{payload, message, status}
}

func SendSuccessJsonResponse(w http.ResponseWriter, payload any) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(NewJsonResponse(payload, "", true))
	if err != nil {
		panic(err)
	}
}

func SendForbiddenResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusForbidden)
	err := json.NewEncoder(w).Encode(NewJsonResponse("", message, true))
	if err != nil {
		panic(err)
	}
}

func SendErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(NewJsonResponse("", message, true))
	if err != nil {
		panic(err)
	}
}
