FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wss-server ./cmd/server/main.go

# Create a minimal production image
FROM alpine:latest  

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/wss-server .

# Expose the application port
EXPOSE 3000

# Run the binary
CMD ["./wss-server"] 