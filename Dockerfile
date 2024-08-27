# STEP 1: Build Phase
# Use an official Golang image as the builder
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies (cached if go.mod and go.sum have not changed)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app (replace "main.go" with your main Go file if it's different)
RUN go build -o main .

# STEP 2: Runtime Phase
# Use a minimal image for running the binary (Alpine in this case)
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the binary from the build phase
COPY --from=builder /app/main .

# Expose the port on which the application will listen (e.g., 8080)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
