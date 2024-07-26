package record

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

func (gl *GameList) GetGame(roomId string) *GameMemory {
	for _, game := range gl.GameList {
		if game.RoomId == roomId {
			return &game
		}
	}

	return nil
}
