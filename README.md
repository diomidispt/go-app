# go-app

A monolith REST API written in Go — pharmaceutical management system for medicines, patients, and prescriptions.

## Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go |
| Database | PostgreSQL 16 |
| Frontend | Vanilla HTML/JS (served by Go) |
| Router | chi |
| Migrations | golang-migrate |
| Local | docker-compose |
| Cloud | ECS Fargate + RDS + ALB (Phase 2) |

## Run locally

**Requirements:** Docker + Docker Compose

```bash
cp .env.example .env       # copy env file (edit values if needed)
docker compose up --build  # starts app + postgres + pgadmin
```

| Service | URL |
|---------|-----|
| App | http://localhost:8080 |
| pgAdmin | http://localhost:5050 |

Migrations run automatically on startup.

## API

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check — pings the database |
| GET/POST | `/medicines` | List or create medicines |
| GET/PUT/DELETE | `/medicines/{id}` | Get, update, or soft delete a medicine |
| GET/POST | `/patients` | List or create patients |
| GET/PUT/DELETE | `/patients/{id}` | Get, update, or soft delete a patient |
| GET/POST | `/prescriptions` | List or create prescriptions |
| GET/PUT/DELETE | `/prescriptions/{id}` | Get, update, or soft delete a prescription |

A full Postman collection is available at `postman/pharma-api.json`.

**Business rule:** `POST /prescriptions` returns `422` if the medicine has zero stock.

## Project structure

```
cmd/api/         — entry point (main.go)
internal/
  app/           — dependency injection (wires all layers)
  config/        — database connection + migrations
  handler/       — HTTP layer (routes, request/response)
  service/       — business logic
  repository/    — SQL queries
  model/         — Go structs (mirrors DB tables)
migrations/      — SQL migration files (golang-migrate)
frontend/        — HTML/CSS/JS frontend
postman/         — Postman collection
```

## Environment variables

| Variable | Description |
|----------|-------------|
| `PORT` | HTTP port (default: 8080) |
| `DB_HOST` | Postgres host |
| `DB_PORT` | Postgres port (default: 5432) |
| `DB_USER` | Postgres user |
| `DB_PASSWORD` | Postgres password |
| `DB_NAME` | Postgres database name |

Locally these come from `.env`. In AWS they are injected from Secrets Manager into ECS.
