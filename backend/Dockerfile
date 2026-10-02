# ==========================================

# Stage 1: Build

# ==========================================

FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend main.go

# ==========================================

# Stage 2: Runtime

# ==========================================

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]

