package v1

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/constants"
	"github.com/robert-w/go-server/internal/routes/v1/sample"
	"go.opentelemetry.io/otel"
)

func RegisterRoutes(subrouter *mux.Router) {
	tracer := otel.Tracer(constants.SERVICE_NAME)

	sampleHandler := &sample.Handler{ Tracer: tracer }
	subrouter.HandleFunc("/samples", sampleHandler.ListSamples).Methods("GET")
	subrouter.HandleFunc("/samples", sampleHandler.CreateSamples).Methods("POST")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.ReadSample).Methods("GET")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.UpdateSample).Methods("PUT")
	subrouter.HandleFunc("/samples/{id}", sampleHandler.DeleteSample).Methods("DELETE")
}
