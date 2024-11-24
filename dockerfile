# Use the official Go image as the build environment
FROM golang:1.22.5-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files (from the root directory)
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire project to the container
COPY . .

# Build the Go app from 'cmd/api'
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Start a new stage from scratch
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose the port where your server will run
EXPOSE 4040

# Command to run the executable
CMD ["./main"]
