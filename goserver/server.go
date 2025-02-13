package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "Chat/goserver/generated" // Make sure this matches your Go module path

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) SendMessage(stream pb.ChatService_SendMessageServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("Received message from client: %s\n", req.GetContent())

		// Simulate typing a response interactively on the server side
		fmt.Print("Server: Type your response: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		responseContent := scanner.Text()

		// Send the typed response back to the client immediately
		serverResponse := &pb.MessageResponse{
			Status:    "RESPONSE FROM SERVER",
			Timestamp: time.Now().Format("2006-01-02 15:04:05") + " - " + responseContent + ". Type your response below",
		}
		if err := stream.Send(serverResponse); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Println("Server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
