package test

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/robert-w/go-server/internal/database"
)

func TestMain(m *testing.M) {
	var err error
	var pool *pgxpool.Pool

	ctx := context.Background()
	// Establish connection to the database
	pool, err = database.NewPool(ctx)
	if err != nil {
		os.Exit(1)
	}

	// Run some setup code
	err = database.RunUpMigration(ctx, pool)
	if err != nil {
		pool.Close()
		os.Exit(1)
	}

	exitCode := m.Run()

	// Clean up from the tests
	err = database.RunDownMigration(ctx, pool)
	if err != nil {
		pool.Close()
		os.Exit(1)
	}

	pool.Close()
	os.Exit(exitCode)
}
