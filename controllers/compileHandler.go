package controllers

import (
	"encoding/json"
	"net/http"

	dockersandbox "remote-code-compiler/dockersandbox"
)

type CodeResponse struct {
	Status string
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

	go dockersandbox.Run(sandbox)

	data := CodeResponse{"Code Submitted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
