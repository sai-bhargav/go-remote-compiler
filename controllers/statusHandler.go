package controllers

import (
	"encoding/json"
	"net/http"
	"remote-code-compiler/services"

	"github.com/gorilla/mux"
)

func StatusHandler(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	statusResponse := services.ExecutionStatus(params["id"])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(statusResponse)
}
