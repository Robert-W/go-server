package server

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robert-w/go-server/internal/requestutil"
	"github.com/robert-w/go-server/internal/routes/system"
	"github.com/robert-w/go-server/internal/routes/v1/user"
)

func registerSystemRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/health", system.Healthcheck).Methods("GET")
}

func registerV1Routes(subrouter *mux.Router, utils *requestutil.RequestUtils, pool *pgxpool.Pool) {
	user.RegisterRoutes(subrouter, utils, pool)
}
