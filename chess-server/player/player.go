package player

import (
	"bytes"
	"kzinthant-d3v/go-chess-server/room"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	Id       string
	conn     *websocket.Conn
	gameRoom *room.Room
	color    string
	Send     chan []byte
}

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (cp *Player) readPump() {
	defer func() {
		cp.gameRoom.leave <- cp
		cp.conn.Close()
	}()
	cp.conn.SetReadLimit(int64(maxMessageSize))
	cp.conn.SetReadDeadline(time.Now().Add(pongWait))
	cp.conn.SetPongHandler(func(string) error { cp.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := cp.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		cp.gameRoom.message <- message
	}
}

func (cp *Player) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		cp.conn.Close()
	}()
	for {
		select {
		case message, ok := <-cp.send:
			cp.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				cp.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := cp.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(cp.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-cp.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			cp.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := cp.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeChessPlayerWs(id string, color string, gameRoom *GameRoom, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	player := &Player{Id: id, color: color, gameRoom: gameRoom, conn: conn, send: make(chan []byte, 256)}

	player.gameRoom.join <- player

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go player.writePump()
	go player.readPump()
}