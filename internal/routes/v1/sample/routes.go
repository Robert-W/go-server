package sample

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/db/services"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func RegisterRoutes(subrouter *mux.Router, tracer oteltrace.Tracer) {
	sampleHandler := &Handler{Service: &services.SampleService{}, Tracer: tracer}
	subrouter.HandleFunc("/samples", sampleHandler.ListSamples).Methods("GET")
	subrouter.HandleFunc("/samples", sampleHandler.CreateSamples).Methods("POST")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.ReadSample).Methods("GET")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.UpdateSample).Methods("PUT")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.DeleteSample).Methods("DELETE")
}
