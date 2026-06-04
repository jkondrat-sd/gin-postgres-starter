# Example Gin

> Why: this README is the starter guide for understanding, running, and testing the project.
> What to do: update it whenever routes, setup commands, or project structure change.

A reusable Go/Gin prototype API using the layered structure from the Notion cheatsheet:

- `handler`: HTTP request/response layer
- `routes`: Gin route grouping and middleware wiring
- `service`: business logic
- `repository`: database access
- `model`: GORM models
- `dto`: request and response structs
- `queue`: Redis-backed background job publishing
- `event`: Kafka-backed domain event publishing
- `worker`: background job handlers
- `pkg`: shared utilities such as database, response, and logging

## Features

- Gin router with `/api/v1`
- PostgreSQL connection through GORM
- User registration and login
- JWT middleware
- Project CRUD for authenticated users
- Health endpoint
- Structured JSON responses
- Environment-based config
- Docker Compose for local PostgreSQL, Redis, and Kafka
- Redis/Asynq background jobs
- Kafka domain events
- Separate worker process for async tasks
- Separate consumer process for event streaming examples
- SQL migration file for reference

## Run Locally

1. Start PostgreSQL, Redis, and Kafka:

```bash
docker compose up -d postgres redis kafka
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the API:

```bash
go run cmd/server/main.go
```

The API runs on `http://localhost:8080`.

4. In another terminal, run the worker:

```bash
go run cmd/worker/main.go
```

When a user registers, the API creates the user and enqueues a `email:send_welcome` job. The worker consumes that job and logs the welcome-email payload.

5. In another terminal, run the Kafka consumer:

```bash
go run cmd/consumer/main.go
```

When a user registers, the API also publishes a `user.registered` event to Kafka topic `user.events`. The consumer reads that event and logs the message.

## Quick Test

```bash
curl http://localhost:8080/health
```

Register:

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"nat@example.com","password":"password123","name":"Nat"}'
```

Login:

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"nat@example.com","password":"password123"}'
```

Create a project:

```bash
curl -X POST http://localhost:8080/api/v1/projects \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"name":"Prototype API","key":"PROTO","description":"Starter project for future apps"}'
```

## Routes

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| GET | `/health` | No | Service health |
| POST | `/api/v1/auth/register` | No | Create user |
| POST | `/api/v1/auth/login` | No | Login and receive JWT |
| GET | `/api/v1/me` | Yes | Current user |
| GET | `/api/v1/projects` | Yes | List projects |
| POST | `/api/v1/projects` | Yes | Create project |
| GET | `/api/v1/projects/:id` | Yes | Get project |
| PUT | `/api/v1/projects/:id` | Yes | Update project |
| DELETE | `/api/v1/projects/:id` | Yes | Delete project |

## Notes

This prototype uses `AutoMigrate` on startup for fast local development. For production projects, prefer explicit migrations from the `migrations` directory.
