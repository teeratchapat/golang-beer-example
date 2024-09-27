# Use the official Golang image as a build stage
FROM golang:1.19 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main cmd/main.go local

# Use a smaller image for the runtime
FROM alpine:latest

WORKDIR /app

# Install MariaDB and MongoDB clients
RUN apk add --no-cache mariadb-client mongodb-tools

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Start the Go application
CMD ["./main"]
