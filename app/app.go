package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	m := mux.NewRouter()

	m.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", m))
}
