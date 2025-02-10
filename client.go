package main

import (
	"context"
	"log"
	"time"

	pb "Chat/generated" // Use relative import path if in the same folder structure

	"google.golang.org/grpc"
)

func main() {
	// Connect to the Python server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	// Send message to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendMessage(ctx, &pb.MessageRequest{Content: "Hello from Go client"})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	log.Printf("Server response: Status: %s, Timestamp: %s", r.GetStatus(), r.GetTimestamp())
}
