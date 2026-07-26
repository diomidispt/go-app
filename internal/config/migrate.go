package config

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/golang-migrate/migrate/v4"                  // the migrate library
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"       // lets migrate read SQL files from disk
)

// RunMigrations applies all pending migrations from the migrations/ folder
// called once on startup before the HTTP server starts
func RunMigrations() error {
	// build the database URL from environment variables — same credentials as the DB connection.
	// Use net/url so special characters in the (AWS-generated) password are encoded in the
	// userinfo; a raw fmt.Sprintf breaks when the password contains ':' '/' '@' etc.
	dbURL := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Host:     net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		Path:     "/" + os.Getenv("DB_NAME"),
		RawQuery: "sslmode=require",
	}).String()

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
