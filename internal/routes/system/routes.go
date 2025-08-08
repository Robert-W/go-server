package system

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/health", healthcheck).Methods("GET")
}

func healthcheck(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}
