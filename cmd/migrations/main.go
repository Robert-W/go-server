package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

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

	// TODO: Do we want to wrap this with a context.WithTimeout
	ctx := context.Background()

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
		os.Exit(2)
	}

	if *operation == "create" && *name == "" {
		slog.Error("Missing required argument. When the 'op' flag is 'create', a 'name' flag is required.")
		os.Exit(2)
	}

	// Creating migrations doesnt require a db or provider, so handle that here
	if *operation == "create" {
		slog.Info("Creating migration")
		if createMigration(*name) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	pool, err = postgres.NewPool(ctx)
	if err != nil {
		slog.Error("Unable to create pool", "pool_error", err)
		os.Exit(1)
	}

	provider, err = goose.NewProvider(
		database.DialectPostgres,
		stdlib.OpenDBFromPool(pool),
		nil,
	)
	if err != nil {
		slog.Error("Unable to create Provider", "provider_error", err)
		cleanup(pool, nil)
		os.Exit(1)
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
		cleanup(pool, provider)
		os.Exit(1)
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
	cleanup(pool, provider)
	os.Exit(0)
}

// Perform any cleanup and close any resources
func cleanup(pool *pgxpool.Pool, provider *goose.Provider) {
	if pool != nil {
		pool.Close()
	}

	if provider != nil {
		provider.Close()
	}
}

// Function to wrap the creation of a migration and the conversion of it's name
// from using a timestamp to using versions
// The return value indicates whether or not these operations were successful
func createMigration(name string) bool {
	err := goose.Create(nil, "migrations", name, "go")
	if err != nil {
		slog.Error("Error creating migration", "error", err)
		return false
	}
	// Convert the migration from a timestamp to a version number
	err = goose.Fix("migrations")
	if err != nil {
		slog.Error("Error fixing migration", "error", err)
		return false
	}

	return true
}
