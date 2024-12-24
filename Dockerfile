# Start with the official Golang image for building the application
FROM golang:1.23.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and dependencies first to take advantage of Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project into the container
COPY . .

# Build the Go app (targeting the main.go file inside cmd/web/)
RUN go build -o gonas ./cmd/web

# Start with a minimal Alpine image to run the app
FROM alpine:latest

# Install necessary libraries for the app to run (e.g., libc)
RUN apk --no-cache add libc6-compat

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder image
COPY --from=builder /app/gonas .

# Expose the port that your app will run on
EXPOSE 8080

# Run the app
CMD ["./gonas"]
