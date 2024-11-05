# Use an official Go runtime as a parent image
FROM golang:1.23.2-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go project inside the container
RUN go build -o sample-go-rest-api ./cmd/server

# Use a minimal base image to copy the binary from the builder stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder environment
COPY --from=builder /app/sample-go-rest-api .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["/root/sample-go-rest-api"]
