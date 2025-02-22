To Dockerize and run the python client and server, few changes need to be made.
Since we need a modular approach, we can not use docker compose as that would lead to all 4 services running together.

We need to run the server with an interactive terminal, then the client. and this needs to be run within a docker network.

First step to do is the changes when running on docker compared to localhost.

in client.py, the hostname needs to be changed from localhost to the server container name, in this case python-server

When running the build and run commands, thy need to be run from within respective folders such as pyserver or pyclient. Running from root will interfere with docker containerization.

after building images using these commands:

PyServer:
'''docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/python-server --push -f pyserver.Dockerfile .'''

GoClient:
'''docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/go-client --push -f goclient.Dockerfile .'''

GoServer:
'''docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/go-server --push .'''

PyClient: 
'''docker buildx build --platform linux/amd64,linux/arm64 -t sufyaansaeed/python-client --push -f pyclient.Dockerfile .'''

After Image creation docker network needs to be created:
'''docker network create chat-network'''

After network creation, you can now run the python server and client with specific commands

PyServer(This command opens the container server terminal):
'''docker run -it --name server --network chat-network -p 50051:50051 sufyaansaeed/python-server'''

GoClient:
'''docker run -it --name go-client --network chat-network sufyaansaeed/go-client'''

GoServer:
'''docker run -it --name server --network chat-network -p 50051:50051 sufyaansaeed/go-server '''

PythonServer:
'''docker run -it --name python-client --network chat-network sufyaansaeed/python-client'''\



Make sure no prio images exist prior to creating and running servers using:
'''docker rm python-server'''