package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connection wraps the websocket connection.
type Connection struct {
	*websocket.Conn
}

// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
func Upgrade(w http.ResponseWriter, r *http.Request) (*Connection, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &Connection{conn}, nil
}
