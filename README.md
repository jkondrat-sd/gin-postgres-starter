# Example Gin

A reusable Go/Gin prototype API using the layered structure from the Notion cheatsheet:

- `handler`: HTTP request/response layer
- `service`: business logic
- `repository`: database access
- `model`: GORM models
- `dto`: request and response structs
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
- Docker Compose for local PostgreSQL
- SQL migration file for reference

## Run Locally

1. Start PostgreSQL:

```bash
docker compose up -d postgres
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
