# Use the Go image as the base for building the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o main .

# Command to run the application
CMD ["./main"]
