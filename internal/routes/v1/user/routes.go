package user

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robert-w/go-server/internal/middleware"
	"github.com/robert-w/go-server/internal/requestutil"
)

func RegisterRoutes(subrouter *mux.Router, utils *requestutil.RequestUtils, pool *pgxpool.Pool) {
	userHandler := &handler{service: &userService{ pool: pool }}
	subrouter.HandleFunc("/users", userHandler.list).Methods("GET")
	subrouter.HandleFunc("/users", middleware.ValidateMiddleware(
		utils,
		&UserPost{},
		userHandler.create,
	)).Methods("POST")

	subrouter.HandleFunc("/users/{id}", userHandler.get).Methods("GET")
	subrouter.HandleFunc("/users/{id}", userHandler.delete).Methods("DELETE")
	subrouter.HandleFunc("/users/{id}", middleware.ValidateMiddleware(
		utils,
		&UserPut{},
		userHandler.update,
	)).Methods("PUT")
}
