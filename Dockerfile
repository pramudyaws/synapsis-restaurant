# Use the Go image as the base for building the application
FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Use Alpine Linux for the final runtime image
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to allow external access
EXPOSE 8080

# Command to run the application
CMD ["./main"]
