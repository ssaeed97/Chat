# Python client Dockerfile

# Use the official Python image as a base
FROM python:3.9-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . /app

# Install necessary dependencies
RUN pip install --no-cache-dir grpcio grpcio-tools

# Generate the gRPC stubs from the proto file
RUN python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. chat.proto

# Command to run the Python client
CMD ["python", "client.py"]
