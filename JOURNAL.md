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

---

## 26/05/2026

### What we built today

**Service layer — all three entities**
`internal/service/medicine.go`, `internal/service/patient.go`, `internal/service/prescription.go`
The business logic layer. Medicines and patients pass straight through to the repository. Prescriptions contain the key business rule: before creating a prescription, the service checks the medicine's stock. If stock is 0 it returns an error and the repository is never called.

**Handler layer — all three entities**
`internal/handler/medicine.go`, `internal/handler/patient.go`, `internal/handler/prescription.go`
The HTTP layer. Each handler reads input from the URL or request body, calls the service, and returns the correct HTTP status code. Prescription errors from business logic return `422 Unprocessable Entity` — not `500` (server crash) and not `400` (bad input).

**`internal/app/app.go` — dependency injection**
Refactored wiring out of `main.go` into a dedicated `App` struct. All repositories, services, and handlers are created here and connected in the correct order. `main.go` now only loads env, connects to DB, calls `app.New(db)`, and starts the server. Adding new entities in future only requires changes to `app.go`.

**Soft delete migrations**
Three additional migrations added:
- `000004_add_deleted_at_to_medicines`
- `000005_add_deleted_at_to_patients`
- `000006_add_deleted_at_to_prescriptions`

Hard delete replaced with soft delete across all entities. All queries filter `WHERE deleted_at IS NULL` so deleted records are invisible to the API but remain in the database for audit purposes.

**`postman/pharma-api.json`**
Single Postman collection with all 16 endpoints organised in four folders: Health, Medicines, Patients, Prescriptions. Committed to the repo so anyone who clones it can test immediately.

### What is working
- Full CRUD REST API for medicines, patients, prescriptions
- Soft delete on all three entities
- Stock check business logic — `POST /prescriptions` returns `422` if medicine stock is 0
- All 16 endpoints tested and verified in Postman

### What is not done yet
- `Dockerfile` — app is not containerised yet
- `docker-compose` — app service not wired in yet
- `frontend/` — no HTML pages yet

### Next
- `Dockerfile` — containerise the Go app
- Wire app service into `docker-compose.yml`
- Run full stack with one command

---

## 26/05/2026 (continued)

### What we built today (Phase 1 completion)

**`Dockerfile`**
Multi-stage build — Stage 1 uses the full Go image to compile the binary, Stage 2 uses Alpine (~20MB) and copies only the binary and migrations. Runs as a non-root user (`appuser`) for security. `CGO_ENABLED=0` produces a statically linked binary that works in a minimal image.

**`docker-compose.yml` — app service wired in**
The `app` service was already scaffolded. Now that the Dockerfile exists, `docker compose up --build` builds and runs all three containers together: `pharma_app`, `pharma_db`, `pharma_pgadmin`. `DB_HOST=db` uses Docker's internal DNS to resolve the Postgres container by name.

**`internal/config/migrate.go` — auto migrations**
`RunMigrations()` runs all pending migrations on startup using the `golang-migrate` library — no manual `migrate up` command needed. Called in `main.go` after DB connect, before the HTTP server starts. `ErrNoChange` is ignored — it means all migrations already ran, which is normal on every restart.

**`frontend/` — B2B UI**
Five files: `index.html`, `medicines.html`, `patients.html`, `prescriptions.html`, `style.css`.
- Homepage: hero section, live stat cards (counts from API), quick action links
- Section pages: labeled forms, tables with stock badges (green/red), confirmation dialogs on delete
- All data comes from the REST API via JavaScript `fetch()` — no page reloads
- Go serves the files as static assets from the `frontend/` directory

### What is working — full Phase 1
- `docker compose up` starts the entire stack with one command
- Migrations run automatically on every startup
- All 16 API endpoints working
- Frontend accessible at http://localhost:8080
- Stock check enforced — UI shows the error message from the API

### Next — Phase 2 (AWS)
- Provision infrastructure in `terraform-sauron`: VPC, RDS, ECR, ECS Fargate, ALB
- GitHub Actions CI/CD: build Docker image → push to ECR → deploy to ECS
- App live on a public URL
