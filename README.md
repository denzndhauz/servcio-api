# Ecommerce API

This is the backend API for an Servcio App.

## Structure

- `cmd/server`: Entry point of the application.
- `internal/auth`: JWT creation, parsing, and middleware.
- `internal/config`: Environment variables and configuration.
- `internal/database`: Database connection and migrations.
- `internal/models`: Database models (User, Product, Order).
- `internal/repository`: Database operations (CRUD).
- `internal/service`: Business logic (ordering, authentication, etc.).
- `internal/handler`: HTTP handlers (controllers).
- `internal/middleware`: HTTP middlewares (auth, logging, etc.).
- `pkg/utils`: Helper functions (password hashing, etc.).

## Getting Started

1. Copy `.env.example` to `.env` and update the values.
2. Run `go mod tidy` to install dependencies.
3. Start the server: `go run ./cmd/server`.
