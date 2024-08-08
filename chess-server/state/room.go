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
	players   [2]*Player
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
		players: [2]*Player{},
		message: make(chan []byte),
		join:    make(chan *Player),
		leave:   make(chan *Player),
		stop:    make(chan struct{}),
		mu:      sync.Mutex{},
	}
}

func (r *Game) Run(ctx context.Context) {
	fmt.Printf("Game started %v\n", r.id)
	r.isRunning = true

	for {
		select {
		case player := <-r.join:
			fmt.Println("someone is trying to join here")
			if r.players[0] != nil && r.players[1] != nil {
				player.send <- []byte("Game is full")
				break
			}

			// Check if the player is already in the game
			for _, existingPlayer := range r.players {
				if existingPlayer != nil && existingPlayer.Id == player.Id {
					player.send <- []byte("You are already in the game")
					return
				}
			}

			// Add the player to the first available slot
			for i, slot := range r.players {
				if slot == nil {
					r.players[i] = player
					break
				}
			}

			fmt.Printf("Player joined the game %v: player id %v\n", r.id, player.Id)

		case player := <-r.leave:
			for i, p := range r.players {
				if p != nil && p.Id == player.Id {
					r.players[i] = nil
					close(player.send)
					fmt.Printf("Player leave the game %v: player id %v\n", r.id, player.Id)

					if r.players[0] == nil && r.players[1] == nil {
						fmt.Println("All players left the game, closing the game")
						close(r.stop)
						ctx.Value(middleware.GameListKey).(*record.GameList).UpdateGameRunning(r.id, false)
					}
					break
				}
			}

		case msg := <-r.message:
			for _, player := range r.players {
				select {
				case player.send <- msg:
					fmt.Printf("Message sent to player %v in game %v, %v\n", player.Id, r.id, msg)
				default:
					close(player.send)
				}
			}

		case <-r.stop:
			fmt.Printf("Game stopped %v\n", r.id)
			r.isRunning = false
			return
		}
	}
}

func (r *Game) CheckRunning() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isRunning
}
