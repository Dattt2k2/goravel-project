# Start from the official Go image
FROM golang:1.22 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o goravel-app

# Create a lightweight production image
FROM alpine:latest

# Add necessary tools
RUN apk --no-cache add ca-certificates tzdata netcat-openbsd

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/goravel-app .
COPY --from=builder /app/.env .
COPY docker-entrypoint.sh .

# Make the entry script executable
RUN chmod +x docker-entrypoint.sh

# Expose the application port
EXPOSE 8000

# Run the application with our entry script
CMD ["./docker-entrypoint.sh"]
