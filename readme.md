# gRPC Multi-Language Chat System

This repository demonstrates a bidirectional chat application using gRPC with implementations in both Python and Go. It includes:

- A **Python server** that handles incoming chat messages and prompts the operator for interactive responses.
- A **Python client** that connects to the Python server and exchanges messages.
- A **Go client** that connects to the Python server and exchanges messages.
- A **Go server** (optional) that can be used as an alternative to the Python server.

All components have been dockerized with Dockerfiles and support multi-architecture builds (both AMD64 and ARM64). A Docker Compose file is also provided to orchestrate the containers easily.

## Overview

The purpose of this project is to showcase:

- Building gRPC services in both Python and Go.
- Implementing bidirectional streaming for a chat system.
- Dockerizing components to run in isolated containers.
- Building multi-architecture Docker images using Docker Buildx.
- Running containers individually or via Docker Compose.

## Project Structure

The directory sructure follows the same template for the 4 functionalities. Each functionality has its own proto file and Dockerfiles. The python features have their stub files within root whereas for Go functionalities, the go.mod and go.sum files are placed in the root but the stub files are placed in /generated sub-directory

## Build

### Multi-Architecture Docker Builds

Use Docker Buildx to build images that work on both AMD64 and ARM64 and run these commands from within the individual directories

### Python Server:
`docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/python-server --push -f pyserver.Dockerfile .`

### Python Client:
`docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/python-client --push -f pyclient.Dockerfile .`

### Go Server:
`docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/go-server --push .`

### Go Client:
`docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/go-client --push -f goclient.Dockerfile .`

## Running Containers Individually:

First, create a common Docker network so containers can communicate:
### Docker Network:
`docker network create chat-network`

### Python Server:
`docker run -it --name server --network chat-network -p 50051:50051 sufyaansaeed/python-server`

### Python Client:
`docker run -it --name python-client --network chat-network sufyaansaeed/python-client`

### Go Server:
`docker run -it --name server --network chat-network -p 50051:50051 sufyaansaeed/go-server`

### Go Client:
`docker run -it --name go-client --network chat-network sufyaansaeed/go-client`


## Usage

### Python Server
- Waits for chat messages from clients. When a message is received, it prompts for an interactive response (via terminal input).

### Python Client & Go Client
- Connect to the Python server and allow you to send messages interactively.
- **Important:** Ensure that in the client code, the server address is set to `python-server:50051` (or the appropriate container/service name) instead of `localhost`.

### Go Server (Optional)
- An alternative implementation of the server in Go that can be used similarly.

## Exiting Containers

- To exit an interactive container (client or server), type `exit` or press `Ctrl+D`.
- To detach from an interactive container without stopping it, press `Ctrl+P` then `Ctrl+Q`.

## Troubleshooting

### Container Name Conflicts
Remove any conflicting containers with:

`docker rm -f <container_name>`

## License

This project is licensed under the MIT License.

## Acknowledgments

- [gRPC](https://grpc.io/)
- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)
- [Python](https://www.python.org/)
