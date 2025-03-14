package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func attachHttpHandlers() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		listenForRenderRequests(websocket)
	})
}

func listenForRenderRequests(conn *websocket.Conn) {
	for {
		messageType, content, err := conn.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("/render received: %s\n", string(content))

		response := fmt.Sprintf("TODO: Render and return image here.")

		if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
			log.Fatal(err)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
