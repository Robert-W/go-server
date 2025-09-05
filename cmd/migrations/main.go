package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"slices"
	"strings"
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

// Exhaustive list of operations that can be passed in as arguments
var VALID_OPERATIONS = []string{"up", "down"}

func main() {
	// Ignore any error here, we just use this locally for convenience
	_ = dotenv.Load()

	logger.SetDefault()

	// Signals I want to catch for graceful shutdown
	signals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	go func() {
		var err error
		var provider *goose.Provider
		var pool *pgxpool.Pool
		var results []*goose.MigrationResult

		// Parse any command line flags here
		operation := flag.String("op", "", fmt.Sprintf("Type of migration to perform, must be one of: [%s]", strings.Join(VALID_OPERATIONS, " ")))
		flag.Parse()

		// Validate any flags and error if appropriate
		if !slices.Contains(VALID_OPERATIONS, *operation) {
			slog.Error("Invalid operation provided", "expected_one_of", VALID_OPERATIONS, "received", *operation)
			stop()
			return
		}

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

		slog.Info("Starting Migrations")

		switch *operation {
		case "up":
			results, err = provider.Up(ctx)
		case "down":
			results, err = provider.DownTo(ctx, CURRENT_VERSION-1)
		}

		if err != nil {
			slog.Error("Failed to run migrations", "migration_error", err)
			cleanup(provider)
			stop()
		}

		for _, result := range results {
			slog.Info("Migration Result",
				"Direction", result.Direction,
				"Duration", result.Duration.Milliseconds(),
				"Empty", result.Empty,
				"Source", result.Source,
			)
		}

		slog.Info(fmt.Sprintf("Ran %d migration(s).", len(results)))
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
