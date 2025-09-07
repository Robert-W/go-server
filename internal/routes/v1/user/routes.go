package user

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(subrouter *mux.Router) {
	userHandler := &handler{service: &userService{}}
	subrouter.HandleFunc("/users", userHandler.list).Methods("GET")
	subrouter.HandleFunc("/users", userHandler.create).Methods("POST")
	subrouter.HandleFunc("/users/{id}", userHandler.get).Methods("GET")
	subrouter.HandleFunc("/users/{id}", userHandler.update).Methods("PUT")
	subrouter.HandleFunc("/users/{id}", userHandler.delete).Methods("DELETE")
}
