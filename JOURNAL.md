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
