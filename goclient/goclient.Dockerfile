# Go client Dockerfile

# Use the official Golang image as a base
FROM golang:1.23-bookworm

# Set the working directory inside the container
WORKDIR /app

# Copy the Go files into the container
COPY . /app

# Install necessary dependencies
#RUN go mod init Chat - Causing errors
#RUN go install google.golang.org/grpc
#RUN go install google.golang.org/protobuf
#RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
#RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#RUN export PATH="$PATH:$(go env GOPATH)/bin"
#RUN go mod tidy

# Install necessary dependencies for protobuf-compiler and Go plugins
RUN apt-get update && apt-get install -y --no-install-recommends \
    protobuf-compiler \
    && rm -rf /var/lib/apt/lists/*  # Clean up to reduce image size

# Install Go tools for generating gRPC code (use specific versions)
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1


# Ensure that Go binaries are in the PATH
RUN export PATH="$PATH:$(go env GOPATH)/bin"

# Generate Go stubs from the .proto file
RUN protoc --go_out=./goclient --go-grpc_out=./goclient goclient/chat.proto

# Expose the port the Go client will communicate through
EXPOSE 50051

# Command to run the Go client
CMD ["go", "run", "client.go"]
