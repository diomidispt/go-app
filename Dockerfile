FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/api .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/frontend ./frontend

RUN adduser -D appuser
USER appuser

EXPOSE 8080

CMD ["./api"]