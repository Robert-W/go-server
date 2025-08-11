package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/robert-w/go-server/internal/constants"
	"github.com/robert-w/go-server/internal/monitoring"
	"github.com/robert-w/go-server/internal/routes/system"
	v1 "github.com/robert-w/go-server/internal/routes/v1"
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

	router.Use(otelmux.Middleware(constants.SERVICE_NAME))

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
