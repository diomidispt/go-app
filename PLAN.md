# Plan

## Use Case
Pharmaceutical — Medicine & Prescription Management System.

Three entities:
- **Medicines** — name, dosage, stock quantity, price
- **Patients** — name, date of birth, contact info
- **Prescriptions** — links a patient to a medicine, with date and instructions

Business logic: cannot create a prescription if the medicine has zero stock.

---

## Architecture

```
Browser (frontend)
    ↓  HTTP requests (JSON)
Go REST API (backend)
    ↓  SQL queries
PostgreSQL (database)
```

- **Language:** Go
- **Database:** PostgreSQL
- **Frontend:** HTML templates served by the Go app
- **Local:** docker-compose (Go + PostgreSQL + pgAdmin)
- **Cloud:** ECS Fargate (Go app) + RDS (PostgreSQL) + ALB

---

## CRUD Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/medicines` | List all medicines |
| GET | `/medicines/{id}` | Get one medicine |
| POST | `/medicines` | Add a new medicine |
| PUT | `/medicines/{id}` | Update a medicine |
| DELETE | `/medicines/{id}` | Delete a medicine |
| GET | `/patients` | List all patients |
| GET | `/patients/{id}` | Get one patient |
| POST | `/patients` | Add a new patient |
| PUT | `/patients/{id}` | Update a patient |
| DELETE | `/patients/{id}` | Delete a patient |
| GET | `/prescriptions` | List all prescriptions |
| GET | `/prescriptions/{id}` | Get one prescription |
| POST | `/prescriptions` | Create a prescription |
| PUT | `/prescriptions/{id}` | Update a prescription |
| DELETE | `/prescriptions/{id}` | Delete a prescription |

---

## Phase 1 — Local (docker-compose)

- [x] Install Go
- [x] `go mod init` — initialise the module
- [x] Create folder structure (`cmd/`, `internal/`, `migrations/`, `frontend/`)
- [x] `cmd/api/main.go` — HTTP server with `/health` endpoint, configurable port
- [x] `docker-compose.yml` — PostgreSQL + pgAdmin + app (app pending Dockerfile)
- [x] `.env` — local secrets, gitignored
- [x] Connect Go to PostgreSQL (`internal/config/db.go`) + `godotenv` for local env loading
- [x] SQL migrations — `medicines`, `patients`, `prescriptions` tables created and applied
- [x] `internal/model/medicine.go` — Medicine struct
- [x] `internal/repository/medicine.go` — CRUD queries for medicines
- [x] `internal/service/medicine.go` — business logic for medicines
- [x] `internal/handler/medicine.go` — HTTP handlers and routes for medicines
- [x] `internal/model/patient.go` — Patient struct
- [x] `internal/repository/patient.go` — CRUD queries for patients
- [x] `internal/service/patient.go` — business logic for patients
- [x] `internal/handler/patient.go` — HTTP handlers and routes for patients
- [x] `internal/model/prescription.go` — Prescription struct
- [x] `internal/repository/prescription.go` — CRUD queries for prescriptions
- [x] `internal/service/prescription.go` — business logic (stock check before creating)
- [x] `internal/handler/prescription.go` — HTTP handlers and routes for prescriptions
- [x] `internal/app/app.go` — dependency injection, all wiring in one place
- [x] Soft delete on all three entities (migrations 000004, 000005, 000006)
- [x] Wire all routes into app.go + main.go
- [x] Test all endpoints with Postman (`postman/pharma-api.json`)
- [x] `Dockerfile` — multi-stage build, Alpine runtime, non-root user
- [x] Wire app service into `docker-compose.yml`
- [x] Auto migrations on startup (`internal/config/migrate.go`)
- [x] Run full stack with `docker-compose up` (app + db + pgadmin)
- [x] `frontend/` — B2B HTML UI, live stat cards, forms, tables, soft delete with confirm

**Phase 1 complete ✅**

## Phase 2 — Cloud (terraform-sauron + ECS)

- [ ] In `terraform-sauron`: provision VPC, subnets, security groups
- [ ] In `terraform-sauron`: provision RDS (PostgreSQL)
- [ ] In `terraform-sauron`: provision ECR (Docker image registry)
- [ ] In `terraform-sauron`: provision ECS cluster + Fargate service
- [ ] In `terraform-sauron`: provision ALB (load balancer in front of ECS)
- [ ] GitHub Actions CI/CD: build Docker image → push to ECR → deploy to ECS
- [ ] App is live on a public URL

---

## Repositories

| Repo | Purpose |
|------|---------|
| `go-app` | Go REST API + HTML frontend |
| `terraform-sauron` | All AWS infrastructure |

---

## Notes
- Secrets locally come from `.env` (gitignored)
- Secrets in AWS come from Secrets Manager (injected into ECS as env vars)
- `DB_HOST` in docker-compose is `db` (Docker resolves container names automatically)
- `DB_HOST` in ECS will be the RDS endpoint from terraform-sauron output
