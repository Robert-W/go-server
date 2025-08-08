package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/routes/system"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
)

type apiServer struct {
	server *http.Server
}

func New(ctx context.Context) *apiServer {
	router := mux.NewRouter()

	// Create all of our subrouters and then pass them into functions to register
	// all the routes in that subpath
	systemRouter := router.PathPrefix("/system").Subrouter()
	v1Router := router.PathPrefix("/v1").Subrouter()

	system.RegisterRoutes(systemRouter)
	v1.RegisterRoutes(v1Router)

	return &apiServer{
		server: &http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: router,
		},
	}
}

func (api *apiServer) Run() error {
	slog.Info("Starting server", "address", api.server.Addr)

	return api.server.ListenAndServe()
}

func (api *apiServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api.server.Shutdown(ctx)
}
