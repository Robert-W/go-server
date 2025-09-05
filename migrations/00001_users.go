package migrations

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUsers, downUsers)
}

func upUsers(ctx context.Context, tx *sql.Tx) error {
	query := "CREATE TABLE IF NOT EXISTS users (" +
		"id uuid PRIMARY KEY," +
		"name TEXT NOT NULL," +
		"created TIMESTAMP default CURRENT_TIMESTAMP NOT NULL," +
		"last_updated TIMESTAMP default CURRENT_TIMESTAMP NOT NULL)"

	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		println("queryy")
		slog.Error("Error executing query", "query", query, "error", err)
		return err
	}

	// This code is executed when the migration is applied.
	return nil
}

func downUsers(ctx context.Context, tx *sql.Tx) error {
	query := "DROP TABLE IF EXISTS users"
	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		slog.Error("Error executing query", "query", query, "error", err)
		return err
	}

	// This code is executed when the migration is rolled back.
	return nil
}
