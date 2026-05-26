module github.com/diomidispt/go-app

go 1.26.3

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect — pgx dependency, handles .pgpass files
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect — pgx dependency, handles pg_service.conf files
	github.com/jackc/pgx/v5 v5.9.2 // indirect — the actual Postgres driver we installed
	github.com/jackc/puddle/v2 v2.2.2 // indirect — pgx dependency, manages connection pooling
	golang.org/x/sync v0.18.0 // indirect — pgx dependency, Go concurrency utilities
	golang.org/x/text v0.31.0 // indirect — pgx dependency, text encoding utilities
)

require (
	github.com/go-chi/chi/v5 v5.3.0 // indirect
	github.com/golang-migrate/migrate/v4 v4.19.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)
