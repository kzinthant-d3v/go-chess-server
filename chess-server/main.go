package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	RoomId   string `json:"roomId"`
	PlayerId string `json:"playerId"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/create-room", func(w http.ResponseWriter, r *http.Request) {
		var data RequestBody

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		w.WriteHeader(http.StatusOK)
		response := map[string]string{"message": "Room created successfully"}
		json.NewEncoder(w).Encode(response)

	}).Methods("POST")

	fmt.Println("Server started at port 5000")
	http.ListenAndServe(":5000", r)
}
