package common

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, code int, payload any) {
	if payload == nil {
		w.WriteHeader(code)
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func SendErrorResponse(w http.ResponseWriter, code int, errors string) {
	SendJSONResponse(w, code, map[string]any{"error": errors})
}

func SendSuccessResponse(w http.ResponseWriter, code int, payload any) {
	SendJSONResponse(w, code, map[string]any{"data": payload})
}
