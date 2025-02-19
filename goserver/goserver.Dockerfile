# Dockerfile (in the goserver folder)

# 1) Build stage
FROM golang:1.23 AS builder

# Create a working directory
WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into /app
COPY . .

# Switch to the directory where server.go actually lives (goserver/ not server/)
WORKDIR /app/goserver

# Build the server binary as a static binary
RUN CGO_ENABLED=0 go build -o server .

# 2) Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary from the builder stage into the final image
COPY --from=builder /app/goserver/server . 

# Ensure the binary has execute permissions
RUN chmod +x ./server

# Expose the gRPC port
EXPOSE 50051

# Run the server
CMD ["./server"]
