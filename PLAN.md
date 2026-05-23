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
- [ ] Connect Go to PostgreSQL (`internal/config/db.go`)
- [ ] Write SQL migrations — create `medicines`, `patients`, `prescriptions` tables
- [ ] `internal/model/` — define Go structs for each entity
- [ ] `internal/repository/` — database queries (CRUD) for each entity
- [ ] `internal/service/` — business logic (e.g. stock check before prescription)
- [ ] `internal/handler/` — HTTP handlers and routes for each endpoint
- [ ] Test all endpoints with Postman
- [ ] `frontend/` — HTML pages that call the API
- [ ] `Dockerfile` — containerise the Go app
- [ ] Run full stack with `docker-compose up` (app + db + pgadmin)

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
