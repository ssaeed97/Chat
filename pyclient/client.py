import grpc
import time
import chat_pb2
import chat_pb2_grpc

def run():
    # Establish connection to the server
    #channel = grpc.insecure_channel('python-server:50051')
    channel = grpc.insecure_channel('server:50051')
    stub = chat_pb2_grpc.ChatServiceStub(channel)

    # Create a stream
    # This is the correct way to initiate bidirectional streaming in gRPC in Python
    def generate_messages():
        while True:
            message = input("Client: ")

            # Exit condition
            if message == "exit":
                print("Closing chat...")
                return

            # Send the message to the server
            yield chat_pb2.MessageRequest(content=message)

    # Start bidirectional streaming (client sends messages, server sends responses)
    responses = stub.SendMessage(generate_messages())

    # Continuously receive responses from the server
    for response in responses:
        print(f"Server response: {response.timestamp}")

if __name__ == '__main__':
    run()
