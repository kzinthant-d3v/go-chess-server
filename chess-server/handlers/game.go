package handlers

import (
	"encoding/json"
	"kzinthant-d3v/go-chess-server/middleware"
	"kzinthant-d3v/go-chess-server/record"
	"kzinthant-d3v/go-chess-server/state"
	"kzinthant-d3v/go-chess-server/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateGameRequestBody struct {
	PlayerId    string `json:"playerId"`
	PlayerColor string `json:"playerColor"`
	PlayerTime  int64  `json:"playerTime"`
}

var runningGames = state.NewRunningGameList()

func ListGamesHandler(w http.ResponseWriter, r *http.Request) {
	gameList := r.Context().Value(middleware.GameListKey).(*record.GameList)

	jsonGamedata, err := json.Marshal(gameList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonGamedata)

}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	var data CreateGameRequestBody

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	gameID, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newGame := record.NewGameMemory(gameID.String(), record.ChessPlayer{
		PlayerId:      data.PlayerId,
		Color:         data.PlayerColor,
		RemainingTime: data.PlayerTime,
		LastActivity:  time.Now().Unix(),
	}, data.PlayerTime)

	gameList := r.Context().Value(middleware.GameListKey).(*record.GameList)
	gameList.AddGame(*newGame)

	jsonGamedata, err := json.Marshal(gameList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonGamedata)

}

type JoinGameRequestBody struct {
	GameId   string `json:"gameId"`
	PlayerId string `json:"playerId"`
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	var data JoinGameRequestBody

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	gameList := r.Context().Value(middleware.GameListKey).(*record.GameList)
	game, err := gameList.GetGame(data.GameId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if game.Player1.PlayerId == data.PlayerId {
		gameList.UpdatePlayer1Joined(data.GameId, true)
	} else if game.Player2.PlayerId == data.PlayerId {
		gameList.UpdatePlayer2Joined(data.GameId, true)
	} else if game.Player2.PlayerId == "" {
		gameList.UpdatePlayer2(data.GameId, record.ChessPlayer{
			PlayerId:      data.PlayerId,
			Color:         utils.ReversePlayer(game.Player1.Color),
			RemainingTime: game.TotalTime,
			LastActivity:  time.Now().Unix(),
			Joined:        true,
		})
	} else {
		http.Error(w, "game not found", http.StatusBadRequest)
		return
	}

	jsonGamedata, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonGamedata))
}

func RunGame(w http.ResponseWriter, r *http.Request) {
	gameId := r.URL.Query().Get("gameId")
	playerId := r.URL.Query().Get("playerId")

	if gameId == "" || playerId == "" {
		http.Error(w, "gameId and playerId are required", http.StatusBadRequest)
		return
	}

	playerList := r.Context().Value(middleware.PlayerListKey).(*record.PlayerGameList)

	_, exist := (*playerList)[playerId]

	if exist {
		http.Error(w, "Player is already in a game", http.StatusBadRequest)
		return
	}

	gameList := r.Context().Value(middleware.GameListKey).(*record.GameList)
	game, err := gameList.GetGame(gameId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := game.GetPlayerByPlayerId(playerId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !player.Joined {
		http.Error(w, "please join the game first", http.StatusBadRequest)
		return
	}

	if !game.IsRunning {
		gameList.UpdateGameRunning(gameId, true)
		newGame := state.NewGame(game.GameId)
		runningGames.AddRunningGame(gameId, newGame)
	}

	gameRoutines, err := runningGames.GetRunningGame(gameId)
	//	fmt.Println("Game routine", gameRoutines)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	go gameRoutines.Run(r.Context(), runningGames)
	state.ServeChessPlayerWs(playerId, gameRoutines, w, r)
}
