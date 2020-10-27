package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Alexplusm/bazaa/projects/chat-app/server/src/controllers"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()

	http.HandleFunc("/ws", controllers.WebSocketHandler)

	fmt.Println("kek")

	log.Fatal(http.ListenAndServe(*addr, nil))
}
