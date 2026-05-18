package Broadcast_Server

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

var url = "ws://localhost:8080/ws"

func Connect() {
	fmt.Println("connecting server... ")
	// Dial server
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	fmt.Println("Connected to WebSocket server.")

	// Channel for handling user input
	inputChan := make(chan string)

	// Goroutine to read user input
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter message: ")
		for {
			if scanner.Scan() {
				inputChan <- scanner.Text()
			} else {
				log.Println("Failed to read input.")
				break
			}
		}
		close(inputChan)
	}()

	// Goroutine to listen for messages from the server
	go func() {
		for {
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Printf("Error reading from WebSocket server: %v", err)
				break
			}
			fmt.Printf("\nServer: %s\n", string(message))
			fmt.Print("Enter message: ") // Re-prompt for user input
		}
	}()

	// Main loop to send user messages
	for msg := range inputChan {
		err := ws.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
	}
	fmt.Println("Disconnected from WebSocket server.")
}