# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/web

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose the port your API runs on (adjust if needed)
EXPOSE 8080

# Run the application
CMD ["./main"]