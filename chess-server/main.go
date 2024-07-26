package main

import (
	"fmt"
	"kzinthant-d3v/go-chess-server/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create-room", handlers.CreateRoomHandler).Methods("POST")

	fmt.Println("Server started at port 5000")
	http.ListenAndServe(":5000", r)
}
