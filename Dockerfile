# Why: this Dockerfile builds a small runnable API image for deployment.
# What to do: keep build steps here, and add runtime files only when the server needs them.
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/server cmd/server/main.go

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /app/bin/server ./server
COPY .env.example ./.env

EXPOSE 8080
CMD ["./server"]
