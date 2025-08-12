package v1

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/constants"
	"github.com/robert-w/go-server/internal/routes/v1/sample"
	"go.opentelemetry.io/otel"
)

func RegisterRoutes(subrouter *mux.Router) {
	tracer := otel.Tracer(constants.SERVICE_NAME)

	sample.RegisterRoutes(subrouter, tracer)
}
