# Dockerfile
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./internal/cmd/receipt_processor

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
