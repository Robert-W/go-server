package v1

import (
	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/routes/v1/sample"
)

func RegisterRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/samples", sample.ListSamples).Methods("GET")
}
