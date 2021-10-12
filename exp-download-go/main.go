package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func get(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		log.Errorf("Error writing to HTTP Response: %s", err)
	}
}

func download(w http.ResponseWriter, req *http.Request) {

	pathVariables := mux.Vars(req)
	id := pathVariables["id"]

	log.Infof("Retrieving resource %s.", id)

	message := fmt.Sprintf("You requested resource %s from the Go Server.\n", id)
	_, err := fmt.Fprintf(w, message)
	if err != nil {
		log.Errorf("Error writing user request to HTTP Response: %s", err)
	}
}

func configureLog() {
	log.SetFormatter(&log.JSONFormatter{})
}

func configureRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", get)
	r.HandleFunc("/{id}/download.txt", download)

	return r
}

func main() {

	configureLog()

	r := configureRoutes()
	http.Handle("/", r)

	log.Info("Setting up listener on port 8081.")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Errorf("Unable to start HTTP server: %s", err)
	}
}
