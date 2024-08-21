# First stage: Build the Go application
FROM golang:1.22.5-bullseye AS builder

WORKDIR /app

# Set environment variables for the build
ENV GOARCH=arm64
ENV GOOS=linux 

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Verify the Go version
RUN go version

# Download dependencies
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o /bin/app

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Second stage: Create a minimal container with the built binary
FROM scratch

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /bin/app /bin/app

# Copy CA certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose the necessary port
EXPOSE 3000

# Command to run the application
CMD ["/bin/app"]
