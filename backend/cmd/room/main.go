package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket connection failed:", err)
		return
	}

	defer conn.Close()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)
		conn.WriteMessage(messageType, msg)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("Websocket server started")
	http.ListenAndServe(":8080", nil)
}