package handlers

import (
	"encoding/json"
	"kzinthant-d3v/go-chess-server/record"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateRoomRequestBody struct {
	PlayerId    string `json:"playerId"`
	PlayerColor string `json:"playerColor"`
	PlayerTime  int64  `json:"playerTime"`
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateRoomRequestBody

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	roomID, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGame := record.NewGameMemory(roomID.String(), record.ChessPlayer{
		PlayerId:      data.PlayerId,
		Color:         data.PlayerColor,
		RemainingTime: data.PlayerTime,
		LastActivity:  time.Now().Unix(),
	}, data.PlayerTime)

	jsonGamedata, err := json.Marshal(newGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonGamedata)

}

type JoinRoomRequestBody struct {
	RoomId   string `json:"roomId"`
	PlayerId string `json:"playerId"`
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {

}
