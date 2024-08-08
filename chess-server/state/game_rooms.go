package state

import (
	"fmt"
	"sync"
)

type RunningGameList struct {
	runningGames map[string]*Game
	mu           sync.Mutex
}

func NewRunningGameList() *RunningGameList {
	return &RunningGameList{
		runningGames: make(map[string]*Game),
		mu:           sync.Mutex{},
	}
}

func (rgl *RunningGameList) AddRunningGame(gameId string, game *Game) {
	rgl.mu.Lock()
	rgl.runningGames[gameId] = game
	rgl.mu.Unlock()
}

func (rgl *RunningGameList) GetRunningGame(gameId string) (*Game, error) {
	rgl.mu.Lock()
	defer rgl.mu.Unlock()
	if game, ok := rgl.runningGames[gameId]; ok {
		return game, nil
	}
	return nil, fmt.Errorf("game not found")
}
