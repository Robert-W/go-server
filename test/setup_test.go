package test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robert-w/go-server/internal/database"
	"github.com/robert-w/go-server/internal/monitoring"
	"go.opentelemetry.io/otel/sdk/trace"
)

func TestMain(m *testing.M) {
	var err error
	var pool *pgxpool.Pool
	var tracerProvider *trace.TracerProvider

	ctx := context.Background()

	// These are set and/or deployed in docker-compose.test.yaml
	os.Setenv("OTEL_COLLECTOR_URL", "0.0.0.0:4318")
	os.Setenv("DATABASE_URL", "postgres://username:password@0.0.0.0:5432/sweet_potato")

	// Setup tracing
	tracerProvider, err = monitoring.NewTraceProvider(ctx)
	if err != nil {
		slog.Error("Error setting up tracing", "error", err)
		os.Exit(1)
	}

	// Establish connection to the database
	pool, err = database.NewPool(ctx, tracerProvider)
	if err != nil {
		slog.Error("Error connecting to database", "error", err)
		tracerProvider.Shutdown(ctx)
		os.Exit(1)
	}

	// Run some setup code
	err = database.RunUpMigration(ctx, pool)
	if err != nil {
		slog.Error("Error running migrations", "error", err)
		tracerProvider.Shutdown(ctx)
		pool.Close()
		os.Exit(1)
	}

	exitCode := m.Run()

	// Clean up from the tests
	err = database.RunDownMigration(ctx, pool)
	if err != nil {
		slog.Error("Error cleaning up database", "error", err)
		tracerProvider.Shutdown(ctx)
		pool.Close()
		os.Exit(1)
	}

	tracerProvider.Shutdown(ctx)
	pool.Close()
	os.Exit(exitCode)
}
