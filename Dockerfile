# Build stage
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and migrations
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary, config files and migrations
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY .env .

EXPOSE 8080

CMD ["./main"] 