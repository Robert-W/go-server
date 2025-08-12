package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/monitoring"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/sdk/trace"
)

type apiServer struct {
	server        *http.Server
	traceProvider *trace.TracerProvider
}

func New(ctx context.Context) (*apiServer, error) {
	traceProvider, err := monitoring.NewTraceProvider(ctx)
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()

	// Create all of our subrouters and then pass them into functions to register
	// all the routes in that subpath
	systemRouter := router.PathPrefix("/system").Subrouter()
	v1Router := router.PathPrefix("/api/v1").Subrouter()

	// this claims to describe the name of the server, but gets mapped to
	// server.address in the spans which is meant for DNS name or IP
	v1Router.Use(otelmux.Middleware("0.0.0.0"))

	registerSystemRoutes(systemRouter)
	registerV1Routes(v1Router)

	return &apiServer{
		server: &http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: router,
		},
		traceProvider: traceProvider,
	}, nil
}

func (api *apiServer) Run() error {
	slog.Info("Starting server", "address", api.server.Addr)

	return api.server.ListenAndServe()
}

func (api *apiServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api.traceProvider.Shutdown(ctx)
	api.server.Shutdown(ctx)
}
