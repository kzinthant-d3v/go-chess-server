package main

import (
	"kzinthant-d3v/go-chess-server/socket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ws", socket.HandleWebSocket)
	log.Println("Server started on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
