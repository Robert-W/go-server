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
		var result *goose.MigrationResult
		var results []*goose.MigrationResult

		VALID_OPERATIONS := []string{"up", "down", "create"}

		// Parse any command line flags here
		operation := flag.String("op", "", fmt.Sprintf("Type of migration to perform, must be one of: [%s]", strings.Join(VALID_OPERATIONS, " ")))
		name := flag.String("name", "", "When op is 'create', the name of the migration we are creating")
		flag.Parse()

		// Validate any flags and error if appropriate
		if !slices.Contains(VALID_OPERATIONS, *operation) {
			slog.Error("Invalid operation provided", "expected_one_of", VALID_OPERATIONS, "received", *operation)
			stop()
			return
		}

		if *operation == "create" && *name == "" {
			slog.Error("Missing required argument. When the 'op' flag is 'create', a 'name' flag is required.")
			stop()
			return
		}

		// Creating migrations doesnt require a db or provider, so handle that here
		if *operation == "create" {
			slog.Info("Creating migration")
			createMigration(*name)
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

		slog.Info("Starting migrations")

		// create was handled above, the only valid remaining cases are up and down
		switch *operation {
		case "up":
			results, err = provider.Up(ctx)
		case "down":
			// Coerce the result into an slice so I can handle the results the same
			// later on in the code
			result, err = provider.Down(ctx)
			results = append(results, result)
		}

		if err != nil {
			slog.Error("Failed to run migrations", "migration_error", err)
			cleanup(provider)
			stop()
		}

		for _, result := range results {
			slog.Info("Migration result",
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

func createMigration(name string) {
	err := goose.Create(nil, "migrations", name, "go")
	if err != nil {
		slog.Error("Error creating migration", "error", err)
	}
	// Convert the migration from a timestamp to a version number
	err = goose.Fix("migrations")
	if err != nil {
		slog.Error("Error fixing migration", "error", err)
	}
}
