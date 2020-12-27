package controllers

import (
	"encoding/json"
	"net/http"

	dockersandbox "remote-code-compiler/services"
)

type CodeResponse struct {
	Message string
}

func RunCode(w http.ResponseWriter, r *http.Request) {

	data := CodeResponse{}

	out := dockersandbox.Run("ruby", "code.rb", "test_cases.txt")

	data.Message = "Remote code compiler generated output" + out
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
