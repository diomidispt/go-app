# go-app

A monolith REST API written in Go.

Built to learn Go and backend engineering from scratch — starting with CRUD operations and growing into a two-tier and three-tier application deployed on AWS.

## What it does

Pharmaceutical management system with three resources: Medicines, Patients, and Prescriptions. Full CRUD operations on each via a REST API.

## Stack

- **Backend:** Go
- **Database:** PostgreSQL
- **Frontend:** HTML templates served by the Go app
- **Local:** docker-compose
- **Cloud:** ECS Fargate + RDS + ALB (coming later)

## Run locally

```bash
docker-compose up -d db pgadmin
go run cmd/api/main.go
```

API available at `http://localhost:8080`
pgAdmin available at `http://localhost:5050`
