package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/CrimsonKarma44/grpc_hello_world/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	g := pb.NewGreeterClient(conn)
	resp, err := g.SayHello(context.Background(), &pb.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", resp.Message)

	c := pb.NewChatClient(conn)
	stream, err := c.ChatStream(context.Background())
	if err != nil {
		log.Fatalf("could not open stream: %v", err)
	}

	// Receive messages in a goroutine
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server closed the stream")
				return
			}
			if err != nil {
				log.Printf("stream error: %v", err)
				return
			}
			log.Printf("%s says: %s", msg.User, msg.Text)
		}
	}()

	// Send messages
	messages := []string{"Hi there", "How’s it going?", "Bye for now"}
	for _, text := range messages {
		err := stream.Send(&pb.ChatMessage{User: "Client", Text: text})
		if err != nil {
			log.Printf("send error: %v", err)
			return
		}
		time.Sleep(1 * time.Second) // Simulate a chat pace
	}
	stream.CloseSend()          // Tell server we’re done sending
	time.Sleep(1 * time.Second) // Let server finish responding
}
