package handler

import (
	"encoding/json"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	if status != 200 || status != 201 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
