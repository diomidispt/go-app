# Journal

---

## 23/05/2026

### What we built today

**`go.mod`**
The Go equivalent of `package.json` (Node.js) or `pom.xml` (Java). It declares the module name and the Go version. Every Go project needs one — created with `go mod init`.

**`docker-compose.yml`**
Defines three services that run together with one command (`docker-compose up`):
- `db` — PostgreSQL 16, the actual database where all data is stored
- `pgadmin` — browser-based UI to visually explore the database at http://localhost:5050
- `app` — the Go API (not running yet, needs a Dockerfile first)

Passwords and config are never hardcoded — they are read from the `.env` file which is gitignored and never committed.

**`.env`**
Lives only on your laptop. Contains the real passwords and config values. Docker Compose reads this automatically. In AWS this will be replaced by Secrets Manager.

**`cmd/api/main.go`**
The entry point of the Go application. When you run the program, this is where it starts. Right now it does one thing: starts an HTTP server and responds with `"ok"` on `/health`. Port is configurable via the `PORT` environment variable.

**Folder structure**
```
cmd/api/          ← entry point, where the program starts
internal/
  config/         ← will load env vars and connect to the database
  handler/        ← will receive HTTP requests and send responses
  service/        ← will contain business logic
  repository/     ← will contain database queries
  model/          ← will define what Medicine, Patient, Prescription look like
migrations/       ← SQL files that create the database tables
frontend/         ← HTML pages the browser loads
```

### What is working
- PostgreSQL running in Docker
- pgAdmin accessible at http://localhost:5050
- Go server starts and responds to `/health`

### What is not connected yet
- Go does not talk to the database yet
- No tables created yet
- No CRUD endpoints yet

### Next
- Connect Go to PostgreSQL (`internal/config/db.go`)
- Write SQL migrations to create the `medicines`, `patients`, `prescriptions` tables
- Build first endpoint: `GET /medicines`

---

## 25/05/2026

### What we built today

**`internal/config/db.go`**
Connects Go to PostgreSQL. Reads the 5 DB credentials from environment variables, builds a DSN (connection string), opens the connection, and pings the database to confirm it's reachable. Returns the connection so the rest of the app can use it.

**`godotenv` dependency**
Installed `github.com/joho/godotenv` and added one line to `main.go` to automatically load the `.env` file when the app starts. Without this, Go couldn't see the DB credentials when run directly with `go run` (Docker Compose injects them automatically, but the terminal does not).

**`golang-migrate` CLI**
Installed the `migrate` CLI tool to manage database schema changes in a versioned, automated way. Each migration is a numbered pair of SQL files — `up` to apply, `down` to roll back. The tool tracks which migrations have already run in a `schema_migrations` table inside Postgres.

**SQL migrations**
Three migrations created and applied:
- `000001_create_medicines` — `id`, `name`, `dosage`, `stock`, `price`
- `000002_create_patients` — `id`, `name`, `date_of_birth`, `phone`, `email`
- `000003_create_prescriptions` — `id`, `patient_id` (FK), `medicine_id` (FK), `date`, `instructions`

Prescriptions references both patients and medicines via foreign keys — Postgres enforces these at the database level.

**`internal/model/medicine.go`**
A Go struct that mirrors the `medicines` table. Each field maps to a column. JSON tags control how the fields appear in API responses (`id`, `name`, `dosage`, `stock`, `price`).

**`internal/repository/medicine.go`**
All SQL queries for the medicines table: `GetAll`, `GetByID`, `Create`, `Update`, `Delete`. Uses `$1`, `$2` placeholders to prevent SQL injection. Uses `RETURNING id` after INSERT to get the auto-generated ID back from Postgres.

### What is working
- Go connects to PostgreSQL on startup
- Three tables exist in the database: `medicines`, `patients`, `prescriptions`
- Medicine model and repository layer are written

### What is not done yet
- `internal/service/medicine.go` — business logic
- `internal/handler/medicine.go` — HTTP handlers and routes
- Same model/repository/service/handler for patients and prescriptions
- No endpoints wired up yet — Postman can only hit `/health`

### Next
- `internal/service/medicine.go` — business logic layer
- `internal/handler/medicine.go` — wire up HTTP routes
- Test `GET /medicines` in Postman with real data
