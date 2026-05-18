package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/CrimsonKarma44/grpc_note/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNoteClient(conn)
	ctx := context.Background()

	// Create a note
	createResp, err := client.CreateNote(ctx, &pb.NoteRequest{
		Title:   "Test Note",
		Content: "This is a test note",
	})
	if err != nil {
		log.Fatalf("failed to create note: %v", err)
	}
	fmt.Printf("Created note: ID=%s, Title=%s, Content=%s\n", createResp.Id, createResp.Title, createResp.Content)

	// Get the note
	getResp, err := client.GetNote(ctx, &pb.AlterNoteRequest{
		Id: createResp.Id,
	})
	if err != nil {
		log.Fatalf("failed to get note: %v", err)
	}
	fmt.Printf("Got note: ID=%s, Title=%s, Content=%s\n", getResp.Id, getResp.Title, getResp.Content)

	// Update the note
	updateResp, err := client.UpdateNote(ctx, &pb.AlterNoteRequest{
		Id:      createResp.Id,
		Title:   "Updated Note",
		Content: "This is an updated test note",
	})
	if err != nil {
		log.Fatalf("failed to update note: %v", err)
	}
	fmt.Printf("Updated note: StatusCode=%d, Message=%s\n", updateResp.StatusCode, updateResp.Message)

	// Delete the note
	deleteResp, err := client.DeleteNote(ctx, &pb.AlterNoteRequest{
		Id: createResp.Id,
	})
	if err != nil {
		log.Fatalf("failed to delete note: %v", err)
	}
	fmt.Printf("Deleted note: StatusCode=%d, Message=%s\n", deleteResp.StatusCode, deleteResp.Message)

	// Try to get deleted note (should fail)
	_, err = client.GetNote(ctx, &pb.AlterNoteRequest{
		Id: createResp.Id,
	})
	if err != nil {
		fmt.Printf("Get deleted note failed as expected: %v\n", err)
	}
}
