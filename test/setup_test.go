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
		os.Exit(1)
	}

	// Run some setup code
	err = database.RunUpMigration(ctx, pool)
	if err != nil {
		slog.Error("Error running migrations", "error", err)
		pool.Close()
		os.Exit(1)
	}

	exitCode := m.Run()

	// Clean up from the tests
	err = database.RunDownMigration(ctx, pool)
	if err != nil {
		slog.Error("Error cleaning up database", "error", err)
		pool.Close()
		os.Exit(1)
	}

	pool.Close()
	os.Exit(exitCode)
}
