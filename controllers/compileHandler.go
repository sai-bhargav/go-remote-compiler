package controllers

import (
	"encoding/json"
	"net/http"

	dockersandbox "remote-code-compiler/services"
)

type CodeResponse struct {
	Message string
}

type RequestPayload struct {
	Language   string
	Code       string
	TestCases  string
	Identifier string
}

func CompileHandler(w http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	sandbox := dockersandbox.Sandbox{}
	err := decoder.Decode(&sandbox)
	if err != nil {
		panic(err)
	}

	out := dockersandbox.Run(sandbox)

	data := CodeResponse{}
	data.Message = "Remote code compiler generated output" + out
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
