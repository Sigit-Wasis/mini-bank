# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mini-bank ./main.go

# Run Stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/mini-bank .
COPY --from=builder /app/.env.example .env

EXPOSE 8080

CMD ["./mini-bank"]
