# Python server Dockerfile

# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install the necessary dependencies
#RUN pip install --no-cache-dir grpcio grpcio-tools
RUN pip install -r requirements.txt

# Generate the gRPC stubs from the proto file
RUN python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. chat.proto


# Expose the port the server will run on
EXPOSE 50051

# Command to run the server
CMD ["python", "server.py"]
