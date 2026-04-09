# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o clothes-store ./cmd/server

FROM alpine:3.19

WORKDIR /app
ENV GIN_MODE=release \
    PORT=8080

COPY --from=builder /app/clothes-store /app/clothes-store
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["/app/clothes-store"]
