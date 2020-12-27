package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string
}

func Health(w http.ResponseWriter, r *http.Request) {

	data := HealthResponse{}
	data.Message = "Remote code compiler is healthy"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
