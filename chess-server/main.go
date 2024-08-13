package main

import (
	"fmt"
	"kzinthant-d3v/go-chess-server/handlers"
	"kzinthant-d3v/go-chess-server/middleware"
	"kzinthant-d3v/go-chess-server/record"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var gameList = record.NewGameList()
var playerList = make(map[string]string)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.WithGameList(gameList))
	r.Use(middleware.WithPlayerList(&playerList))

	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000", "http://127.0.0.1:3000", "http://0.0.0.0:3000"}),
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root route working"))
	})
	r.HandleFunc("/create-game", handlers.CreateGameHandler).Methods("POST")
	r.HandleFunc("/list-games", handlers.ListGamesHandler).Methods("GET")
	r.HandleFunc("/join-game", handlers.JoinGameHandler).Methods("POST")
	r.HandleFunc("/connect-game", handlers.RunGame)

	fmt.Println("Server started at port 5000")
	http.ListenAndServe(":5000", corsHandler)
}
