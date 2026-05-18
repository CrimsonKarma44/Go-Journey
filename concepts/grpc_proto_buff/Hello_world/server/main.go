package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "io"

    pb "github.com/CrimsonKarma44/grpc_hello_world/proto"
    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreeterServer
    pb.UnimplementedChatServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.Name)}, nil
}


func (s *server) ChatStream(stream pb.Chat_ChatStreamServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            return nil // Client closed the stream
        }
        if err != nil {
            return err
        }
        log.Printf("Received from %s: %s", msg.User, msg.Text)
        // Echo back with a server twist
        err = stream.Send(&pb.ChatMessage{
            User:  "Server",
            Text:  "Echo: " + msg.Text,
        })
        if err != nil {
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
    pb.RegisterGreeterServer(s, &server{})
    pb.RegisterChatServer(s, &server{})
    log.Println("Server running on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}