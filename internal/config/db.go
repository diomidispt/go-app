package config

import (
	"database/sql" // Go's built-in database interface — defines how to talk to any database
	"fmt"          // used to build the connection string by combining variables into one string
	"os"           // used to read DB credentials from environment variables

	_ "github.com/jackc/pgx/v5/stdlib" // the Postgres driver we installed — the _ means we import it only for its side effects (it registers itself with database/sql)
)

// Connect opens a connection to PostgreSQL and returns it
// the rest of the app will call this function to get the database connection
func Connect() (*sql.DB, error) {
	// step 1 — read each credential from the environment variables (set in .env)
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// step 2 — combine all credentials into one connection string (called a DSN — Data Source Name)
	// this is the same concept as a connection string in .NET or a JDBC URL in Java
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// step 3 — open the connection using the pgx driver
	// sql.Open does not actually connect yet — it just prepares the connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// step 4 — ping the database to confirm it is actually reachable
	// this is the step that actually tries to connect
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to reach database: %w", err)
	}

	return db, nil // connection is open and working — return it to whoever called Connect()
}
