package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "goclient/generated" // Make sure this matches your Go module path

	"google.golang.org/grpc"
)

func main() {
	// Establish a connection to the server
	conn, err := grpc.Dial("server:50051", grpc.WithInsecure()) //used with docker run
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //Used with local run
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	// Create a stream
	stream, err := c.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}

	// Start a goroutine to receive messages from the server
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("could not receive message: %v", err)
			}
			fmt.Printf("\nServer response: Status: %s\n, Timestamp: %s\n\n", resp.GetStatus(), resp.GetTimestamp())
		}
	}()

	// Continuously read input from the user and send messages
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your message and press Enter (Type 'exit' to quit):")
	for {
		fmt.Print("Client: ")
		scanner.Scan()
		text := scanner.Text()

		// Exit condition
		if text == "exit" {
			fmt.Println("Closing chat...")
			stream.CloseSend()
			break
		}

		// Send the message to the server
		err := stream.Send(&pb.MessageRequest{Content: text})
		if err != nil {
			log.Fatalf("could not send message: %v", err)
		}

		// Add a small delay between messages to simulate typing delay
		time.Sleep(1 * time.Second)
	}
}
