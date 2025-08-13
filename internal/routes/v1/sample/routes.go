package sample

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(subrouter *mux.Router) {
	sampleHandler := &handler{service: &sampleService{}}
	subrouter.HandleFunc("/samples", sampleHandler.listSamples).Methods("GET")
	subrouter.HandleFunc("/samples", sampleHandler.createSamples).Methods("POST")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.readSample).Methods("GET")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.updateSample).Methods("PUT")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.deleteSample).Methods("DELETE")
}
