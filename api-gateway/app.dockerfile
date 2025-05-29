# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api-gateway ./cmd/api-gateway

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api-gateway .
COPY --from=builder /app/api-gateway/internal/config/config.yml ./internal/config/

# Expose port
EXPOSE 8080

# Run the application
CMD ["/app/api-gateway"]
