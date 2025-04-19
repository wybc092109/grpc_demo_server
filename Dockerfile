# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Set GOPROXY for better dependency download
ENV GOPROXY=https://goproxy.cn,direct

# Copy go mod and sum files
COPY go.mod go.sum ./

# Verify and download dependencies
RUN go mod verify && \
    go mod tidy && \
    go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server server.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/server /app/server
COPY --from=builder /app/etc/server.yaml /app/etc/

EXPOSE 8888

# Run the service
CMD ["/app/server", "-f", "/app/etc/server.yaml"]