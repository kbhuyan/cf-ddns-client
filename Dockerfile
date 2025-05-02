# Use a minimal Go image for building
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Pre-copy dependencies for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cf-ddns-client .

# Use a minimal alpine image for the final container
FROM alpine:latest

WORKDIR /root/

# Copy the static binary from the builder stage
COPY --from=builder /app/cf-ddns-client .

# Optional: Install ca-certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

# Command to run the client
CMD ["./cf-ddns-client"]
