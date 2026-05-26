package config

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"                  // the migrate library
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"       // lets migrate read SQL files from disk
)

// RunMigrations applies all pending migrations from the migrations/ folder
// called once on startup before the HTTP server starts
func RunMigrations() error {
	// build the database URL from environment variables — same credentials as the DB connection
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	m, err := migrate.New("file://migrations", dbURL) // point to the migrations folder
	if err != nil {
		return fmt.Errorf("failed to initialise migrations: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err) // ErrNoChange is not an error — it means all migrations already ran
	}

	return nil
}
