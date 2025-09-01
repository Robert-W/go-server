package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
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

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test that we can connect by pinging the connection
	err = pool.Ping(ctx)

	return pool, err
}
