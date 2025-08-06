package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type apiServer struct {
	server *http.Server
}

func New(ctx context.Context) *apiServer {
	return &apiServer{
		server: &http.Server{
			Addr: "0.0.0.0:3000",
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
