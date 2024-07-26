package record

type ChessPlayer struct {
	PlayerId      string
	Color         string
	RemainingTime int64
	LastActivity  int64
}

type GameMemory struct {
	RoomId     string      `json:"roomId"`
	IsFinished bool        `json:"isFinished"`
	Player1    ChessPlayer `json:"player1"`
	Player2    ChessPlayer `json:"player2"`
	TotalTime  int64       `json:"totalTime"`
}

func NewGameMemory(roomId string, player1 ChessPlayer, totalTime int64) *GameMemory {
	return &GameMemory{
		RoomId:     roomId,
		IsFinished: false,
		Player1:    player1,
		Player2:    ChessPlayer{},
		TotalTime:  totalTime,
	}
}

func (gm *GameMemory) JoinGame(player2 ChessPlayer) {
	gm.Player2 = player2
}
