package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// Client represents the websocket client at the server
type Client struct {
	// Actual websocket connection
	conn *websocket.Conn
}

func createClient(conn *websocket.Conn) *Client {
	return &Client{conn: conn}
}

// WebSocketHandler handle websocket request
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request handled!")

	// TODO: remove
	upgrader.CheckOrigin = func(r *http.Request) bool {
		fmt.Println("CheckOrigin!")
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := createClient(conn)
	fmt.Println("New client joined the hub!")
	fmt.Println(client)
}
