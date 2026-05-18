package main

import (
	// "github.com/CrimsonKarma44/grpc_note/note"
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/CrimsonKarma44/grpc_note/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)


type server struct {
	pb.UnimplementedNoteServer
}
var notes = make(map[string]*pb.NoteResponse)

func (s *server) CreateNote(ctx context.Context, req *pb.NoteRequest) (*pb.NoteResponse, error) {
	id := uuid.New().String() // Replace with actual UUID generation logic
	
	resp := &pb.NoteResponse{
		Id:      id,
		Title:   req.Title,
		Content: req.Content,
	}
	notes[id] = resp
	return resp, nil
}

func (s *server) GetNote(ctx context.Context, req *pb.AlterNoteRequest) (*pb.NoteResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("note id is required")
	}
	note, exists := notes[req.Id]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}
	return note, nil
}


func (s *server) UpdateNote(ctx context.Context, req *pb.AlterNoteRequest) (*pb.AlterNoteResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("note id is required")
	}
	if req.Title == "" && req.Content == "" {
		return nil, fmt.Errorf("note title or content is required")
	}
	// Find the note by id and update it
	note, exists := notes[req.Id]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}
	if req.Title != "" {
		note.Title = req.Title
	}
	if req.Content != "" {
		note.Content = req.Content
	}
	
	notes[req.Id] = note
	
	resp := &pb.AlterNoteResponse{
		StatusCode: 200,
		Message:    "Note updated successfully",
	}
	
	return resp, nil
}

func (s *server) DeleteNote(ctx context.Context, req *pb.AlterNoteRequest) (*pb.AlterNoteResponse, error) {
	if req.Id == "" {
		return nil, fmt.Errorf("note id is required")
	}
	_, exists := notes[req.Id]
	if !exists {
		return nil, fmt.Errorf("note not found")
	}
	delete(notes, req.Id)
	resp := &pb.AlterNoteResponse{
		StatusCode: 200,
		Message:    "Note deleted successfully",
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNoteServer(s, &server{})
	log.Println("Server is listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
