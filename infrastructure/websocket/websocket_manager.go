package websocket

import (
	"log"
	"sync"
)

type WebSocketManager struct {
	clients   map[*Connection]bool
	broadcast chan EventUpdateNotification
	mutex     sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:   make(map[*Connection]bool),
		broadcast: make(chan EventUpdateNotification, 1024),
	}
}

func (m *WebSocketManager) AddWebSocketClient(client *Connection) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clients[client] = true
}

func (m *WebSocketManager) RemoveWebSocketClient(client *Connection) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.clients, client)
}

func (m *WebSocketManager) AddToBroadcast(notification EventUpdateNotification) {
	m.broadcast <- notification
}

func (m *WebSocketManager) BroadcastWebSocketMessage() {
	for notification := range m.broadcast {
		m.mutex.Lock()
		for client := range m.clients {
			err := client.WriteJSON(notification)
			if err != nil {
				log.Printf("Error: %v", err)
				_ = client.Close()
				delete(m.clients, client)
			}
		}
		m.mutex.Unlock()
	}
}
