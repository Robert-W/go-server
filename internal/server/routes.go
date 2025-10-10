package server

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/noir-cats/go-sample-server/internal/requestutil"
	"github.com/noir-cats/go-sample-server/internal/routes/system"
	"github.com/noir-cats/go-sample-server/internal/routes/v1/user"
)

func registerSystemRoutes(subrouter *mux.Router) {
	subrouter.HandleFunc("/health", system.Healthcheck).Methods("GET")
}

func registerV1Routes(subrouter *mux.Router, utils *requestutil.RequestUtils, pool *pgxpool.Pool) {
	user.RegisterRoutes(subrouter, utils, pool)
}
