import grpc
from concurrent import futures
import time
import chat_pb2
import chat_pb2_grpc
import threading

class ChatServicer(chat_pb2_grpc.ChatServiceServicer):
    def SendMessage(self, request_iterator, context):
        for request in request_iterator:
            print(f"Received message from client: {request.content}")

            # Simulate typing a response interactively on the server side
            response_content = input("Server: Type your response: ")

            # Send the typed response back to the client immediately
            server_response = chat_pb2.MessageResponse(
                status="RESPONSE FROM SERVER",
                timestamp=time.strftime("%Y-%m-%d %H:%M:%S", time.gmtime()) + f" - {response_content}.Type your response below"
            )
            yield server_response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServiceServicer_to_server(ChatServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server running on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
