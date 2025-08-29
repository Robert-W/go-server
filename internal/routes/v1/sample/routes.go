package sample

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(subrouter *mux.Router) {
	sampleHandler := &handler{service: &sampleService{}}
	subrouter.HandleFunc("/samples", sampleHandler.list).Methods("GET")
	subrouter.HandleFunc("/samples", sampleHandler.create).Methods("POST")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.get).Methods("GET")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.update).Methods("PUT")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.delete).Methods("DELETE")
}
