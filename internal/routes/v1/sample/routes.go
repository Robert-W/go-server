package sample

import (
	"github.com/gorilla/mux"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func RegisterRoutes(subrouter *mux.Router, tracer oteltrace.Tracer) {
	sampleHandler := &handler{service: &sampleService{}, tracer: tracer}
	subrouter.HandleFunc("/samples", sampleHandler.listSamples).Methods("GET")
	subrouter.HandleFunc("/samples", sampleHandler.createSamples).Methods("POST")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.readSample).Methods("GET")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.updateSample).Methods("PUT")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.deleteSample).Methods("DELETE")
}
