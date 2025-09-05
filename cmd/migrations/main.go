package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	dotenv "github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
	postgres "github.com/robert-w/go-server/internal/database"
	"github.com/robert-w/go-server/internal/logger"
	_ "github.com/robert-w/go-server/migrations"
)

// This is the current migration version. If we are unable to apply the
// migrations, to be safe, we roll back to this version.
const CURRENT_VERSION = 1

func main() {
	// Ignore any error here, we just use this locally for convenience
	_ = dotenv.Load()

	logger.SetDefault()
	slog.Info("Starting Migrations")

	// Signals I want to catch for graceful shutdown
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	go func() {
		var provider *goose.Provider
		var pool *pgxpool.Pool
		var err error

		pool, err = postgres.NewPool(ctx)
		if err != nil {
			slog.Error("Unable to create pool", "pool_error", err)
			cleanup(provider)
			stop()
			return
		}

		provider, err = goose.NewProvider(
			database.DialectPostgres,
			stdlib.OpenDBFromPool(pool),
			nil,
		)
		if err != nil {
			slog.Error("Unable to create Provider", "provider_error", err)
			cleanup(provider)
			stop()
			return
		}

		// Temporary until I figure out how I want to invoke this
		// Command line arguments to run up or down migrations
		// provider.DownTo(ctx, CURRENT_VERSION-1)

		results, err := provider.Up(ctx)
		if err != nil {
			slog.Error("Failed to run migrations", "migration_error", err)
			cleanup(provider)
			stop()
		}

		for _, result := range results {
			slog.Info("Migration Result",
				"Source", result.Source,
				"Direction", result.Direction,
				"Duration", result.Duration.Milliseconds(),
			)
		}

		slog.Info(fmt.Sprintf("Ran %d migrations.", len(results)))
		stop()
	}()

	// Wait for a signal and start shutting things down
	<-ctx.Done()
}

func cleanup(provider *goose.Provider) {
	if provider != nil {
		provider.Close()
	}
}
