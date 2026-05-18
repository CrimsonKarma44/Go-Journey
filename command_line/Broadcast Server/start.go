package Broadcast_Server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func Url() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handler)
	go handleMessages()
	fmt.Println("starting server at :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Establishing websocket connection at :8080/ws ...")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for client := range clients {
		fmt.Println(client.RemoteAddr().String(), "connected!")
	}
	defer conn.Close()
	clients[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String()+":", string(message))
		broadcast <- message
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}