package v1

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/routes/v1/sample"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func RegisterRoutes(tracer oteltrace.Tracer, subrouter *mux.Router) {
	sampleHandler := &sample.Handler{ Tracer: tracer }
	subrouter.HandleFunc("/samples", sampleHandler.ListSamples).Methods("GET")
}
