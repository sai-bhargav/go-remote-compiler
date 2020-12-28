package main

import (
	"log"
	"net/http"
	"remote-code-compiler/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", controllers.HealthHandler).Methods("GET")
	router.HandleFunc("/compile", controllers.CompileHandler).Methods("POST")
	router.HandleFunc("/status/{id}", controllers.StatusHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
