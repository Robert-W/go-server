package database

import (
	"context"
	"os"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewPool(ctx context.Context, tracerProvider *trace.TracerProvider) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Set some sensible defaults, these can always be overriden at a later date
	// May require some testing to figure out the best values here
	config.MaxConnIdleTime = 10 * time.Minute
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConns = 8
	config.MinConns = 4
	config.MinIdleConns = 4

	config.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithTracerProvider(tracerProvider))

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test that we can connect by pinging the connection
	err = pool.Ping(ctx)

	return pool, err
}
