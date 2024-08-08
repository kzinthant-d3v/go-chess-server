package record

import "errors"

type ChessPlayer struct {
	PlayerId      string
	Color         string
	RemainingTime int64
	LastActivity  int64
	Joined        bool
}

type GameMemory struct {
	GameId     string      `json:"gameId"`
	IsFinished bool        `json:"isFinished"`
	Player1    ChessPlayer `json:"player1"`
	Player2    ChessPlayer `json:"player2"`
	TotalTime  int64       `json:"totalTime"`
	IsRunning  bool        `json:"isRunning"`
}

func NewGameMemory(gameId string, player1 ChessPlayer, totalTime int64) *GameMemory {
	return &GameMemory{
		GameId:     gameId,
		IsFinished: false,
		Player1:    player1,
		Player2:    ChessPlayer{},
		TotalTime:  totalTime,
		IsRunning:  false,
	}
}

func (gm *GameMemory) GetPlayerByPlayerId(playerId string) (*ChessPlayer, error) {
	if gm.Player1.PlayerId == playerId {
		return &gm.Player1, nil
	}
	if gm.Player2.PlayerId == playerId {
		return &gm.Player2, nil
	}
	return nil, errors.New("player not found")
}

type GameList struct {
	GameList []GameMemory `json:"gameList"`
}

func NewGameList() *GameList {
	return &GameList{
		GameList: []GameMemory{},
	}
}

func (gl *GameList) AddGame(game GameMemory) {
	gl.GameList = append(gl.GameList, game)
}

func (gl *GameList) GetGame(gameId string) (*GameMemory, error) {
	for _, game := range gl.GameList {
		if game.GameId == gameId {
			return &game, nil
		}
	}

	return nil, errors.New("game not found")
}

func (gl *GameList) UpdatePlayer2(gameId string, player2 ChessPlayer) error {
	for i, game := range gl.GameList {
		if game.GameId == gameId {
			gl.GameList[i].Player2 = player2
			return nil
		}
	}

	return errors.New("game not found")
}

func (gl *GameList) UpdatePlayer1Joined(gameId string, val bool) error {
	for i, game := range gl.GameList {
		if game.GameId == gameId {
			gl.GameList[i].Player1.Joined = val
			return nil
		}
	}
	return errors.New("game not found")
}

func (gl *GameList) UpdatePlayer2Joined(gameId string, val bool) error {
	for i, game := range gl.GameList {
		if game.GameId == gameId {
			gl.GameList[i].Player2.Joined = val
			return nil
		}
	}
	return errors.New("game not found")
}

func (gl *GameList) UpdateGameRunning(gameId string, val bool) error {
	for i, game := range gl.GameList {
		if game.GameId == gameId {
			gl.GameList[i].IsRunning = val
			return nil
		}
	}
	return errors.New("game not found")
}
