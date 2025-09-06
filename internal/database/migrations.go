package database

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
)

func RunUpMigration(ctx context.Context, pool *pgxpool.Pool) error {
	provider, err := goose.NewProvider(
		database.DialectPostgres,
		stdlib.OpenDBFromPool(pool),
		os.DirFS(getMigrationsDirectory()),
	)
	if err != nil {
		return err
	}

	result, err := provider.Up(ctx)
	if err != nil {
		return err
	}

	printResult(result)
	return nil
}

func RunDownMigration(ctx context.Context, pool *pgxpool.Pool) error {
	provider, err := goose.NewProvider(
		database.DialectPostgres,
		stdlib.OpenDBFromPool(pool),
		os.DirFS(getMigrationsDirectory()),
	)
	if err != nil {
		return err
	}

	result, err := provider.Down(ctx)
	if err != nil {
		return err
	}

	printResult([]*goose.MigrationResult{result})
	return nil
}

func printResult(results []*goose.MigrationResult) {
	for _, result := range results {
		slog.Info("Migration result",
			"Direction", result.Direction,
			"Duration", result.Duration.Milliseconds(),
			"Empty", result.Empty,
			"Source", result.Source,
		)
	}
}

// Gets the migrations directory path from the root by traversing from here to
// there, if we move this file/function, the path must be updated
func getMigrationsDirectory() string {
	_, file, _, _ := runtime.Caller(0)
	migrationsDir := filepath.Join(filepath.Dir(file), "../../migrations")

	return migrationsDir
}
