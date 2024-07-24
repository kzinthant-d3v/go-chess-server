package main

import (
	"fmt"
	"sync"
)

type Room struct {
	id        string
	players   [2]*Player
	message   chan []byte
	join      chan *Player
	leave     chan *Player
	isRunning bool
	stop      chan struct{}
	mu        sync.Mutex
}

func NewRoom(id string) *Room {
	return &Room{
		id:      id,
		players: [2]*Player{},
		message: make(chan []byte),
		join:    make(chan *Player),
		leave:   make(chan *Player),
		stop:    make(chan struct{}),
		mu:      sync.Mutex{},
	}
}

func (r *Room) Run() {
	fmt.Printf("Game room started %v\n", r.id)
	r.isRunning = true

	for {
		select {
		case player := <-r.join:
			fmt.Println("someone is trying to join here")
			if r.players[0] != nil && r.players[1] != nil {
				player.send <- []byte("Room is full")
				break
			}

			if r.players[0] == nil {
				r.players[0] = player
			} else {
				r.players[1] = player
			}

			fmt.Printf("Player joined the game room %v: player id %v\n", r.id, player.Id)

		case player := <-r.leave:
			for i, p := range r.players {
				if p != nil && p.Id == player.Id {
					r.players[i] = nil
					close(player.send)
					fmt.Printf("Player leave the game room %v: player id %v\n", r.id, player.Id)

					if r.players[0] == nil && r.players[1] == nil {
						close(r.stop)
					}

					break
				}
			}

		case msg := <-r.message:
			for _, player := range r.players {
				select {
				case player.send <- msg:
					fmt.Printf("Message sent to player %v in game room %v, %v\n", player.Id, r.id, msg)
				default:
					close(player.send)
				}
			}

		case <-r.stop:
			fmt.Printf("Game room stopped %v\n", r.id)
			r.isRunning = false
			return
		}
	}
}

func (r *Room) checkRunning() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isRunning
}
