import grpc
from concurrent import futures
import time
import chat_pb2
import chat_pb2_grpc

class ChatServicer(chat_pb2_grpc.ChatServiceServicer):
    def SendMessage(self, request, context):
        print(f"Received message: {request.content}")
        return chat_pb2.MessageResponse(
            status="DELIVERED",
            timestamp=time.strftime("%Y-%m-%d %H:%M:%S", time.gmtime())
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServiceServicer_to_server(ChatServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server running on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
