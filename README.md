# clean-fiber

Fiber scaffold with clean architecture.

The concepts are aligned with [clean-axum](https://github.com/kigawas/clean-axum).

## Features

- [Fiber](https://github.com/gofiber/fiber) framework
- [Gorm](https://github.com/go-gorm/gorm) domain models
- Completely separated API routers and DB-related logic (named "persistence" layer)
- Distinct input parameters, queries and output schemas

## Module Hierarchy

### API Logic

- `api/routers`: Fiber route handlers
- `api/models`: API-specific models (e.g., error responses)
- `api/init.go`: Router initialization
- `api/validator.go`: JSON parameter validation
- `api/ws`: Fiber websocket middleware

### API-agnostic Application Logic

Main concept: Web framework is replaceable.

All modules here should not include any specific API web framework logic.

- `app/persistence`: Database operations
- `app/database.go`: Application DB connection
- `app/config.go`: Application configuration

### DB-API agnostic Domain Models

Main concept: Database (Sqlite/MySQL/PostgreSQL) is replaceable.

Except `gorm` tags of `models/domains`, all modules are ORM library agnostic.

- `models/domains`: Core domain models
- `models/params`: Input parameters for creating/updating domain models
- `models/schemas`: Output schemas for API responses
- `models/queries`: Queries for filtering domain models

### Tests

- `tests/api`: API integration tests. Hierarchy is the same as `api/routers`
- `tests/app`: DB manipulation unit tests. Hierarchy is the same as `app/persistence`

### Others

- `main.go`: Application entry point

## Run

### Start server

```bash
go run main.go

# or with air for auto reload
# go install github.com/air-verse/air@latest
air

# or build for production
go build
./clean-fiber
```

### Benchmark

```bash
# edit .env to use Postgres
# go build
./clean-fiber
wrk --latency -t20 -c50 -d10s http://localhost:3001/users\?username\=
```

### Run tests

```bash
go test -v ./tests/*
```

### Run lint

```bash
# install: https://golangci-lint.run/welcome/install/
golangci-lint run
```
