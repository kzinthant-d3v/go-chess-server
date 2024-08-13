package state

import (
	"context"
	"fmt"
	"kzinthant-d3v/go-chess-server/middleware"
	"kzinthant-d3v/go-chess-server/record"
	"sync"
)

type Game struct {
	id        string
	players   []*Player // Slice to hold players
	message   chan []byte
	join      chan *Player
	leave     chan *Player
	isRunning bool
	stop      chan struct{}
	mu        sync.Mutex
}

func NewGame(id string) *Game {
	return &Game{
		id:      id,
		players: []*Player{},
		message: make(chan []byte),
		join:    make(chan *Player),
		leave:   make(chan *Player),
		stop:    make(chan struct{}),
		mu:      sync.Mutex{},
	}
}

func (g *Game) addPlayer(player *Player) {
	g.players = append(g.players, player)
}

func (g *Game) removePlayer(remoteAddr string) bool {
	for i, player := range g.players {
		if player != nil && player.conn.RemoteAddr().String() == remoteAddr {
			g.players = append(g.players[:i], g.players[i+1:]...) // Remove the player
			return true
		}
	}
	return false
}

func (g *Game) Run(ctx context.Context, runningGames *RunningGameList) {
	fmt.Printf("Game started %v\n", g.id)
	g.isRunning = true

	for {
		select {
		case player := <-g.join:
			g.mu.Lock()
			g.addPlayer(player)
			fmt.Printf("Player %v joined game %v. Current players: %d", player.Id, g.id, len(g.players))
			g.mu.Unlock()

		case player := <-g.leave:
			close(player.send)
			g.mu.Lock()
			if g.removePlayer(player.conn.RemoteAddr().String()) {
				fmt.Printf("Player %v left game %v. Current players: %d", player.Id, g.id, len(g.players))
				if len(g.players) == 0 {
					close(g.message)
					close(g.stop)
					fmt.Printf("All players have left. Game %v is closing.\n", g.id)
				}
			} else {
				fmt.Printf("Player with connection %v not found in game %v.", player.conn.RemoteAddr().String(), g.id)
			}
			g.mu.Unlock()

		case msg := <-g.message:
			g.mu.Lock()
			for _, player := range g.players {
				select {
				case player.send <- msg:
					fmt.Printf("Message sent to player %v in game %v: %s\n", player.Id, g.id, string(msg))
				default:
					fmt.Printf("Warning: Player %v's channel is full, message dropped in game %v\n", player.Id, g.id)
				}
			}
			g.mu.Unlock()

		case <-g.stop:
			ctx.Value(middleware.GameListKey).(*record.GameList).UpdateGameRunning(g.id, false)
			fmt.Printf("Game stopped %v\n", g.id)
			g.isRunning = false
			runningGames.RemoveRunningGame(g.id)
			return
		}
	}
}

func (r *Game) CheckRunning() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isRunning
}
