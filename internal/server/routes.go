package server

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/routes/system"
	"github.com/robert-w/go-server/internal/routes/v1/sample"
)

func registerSystemRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/health", system.Healthcheck).Methods("GET")
}

func registerV1Routes(subrouter *mux.Router) {
	sample.RegisterRoutes(subrouter)
}
