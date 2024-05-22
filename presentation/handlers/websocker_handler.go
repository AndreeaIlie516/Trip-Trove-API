package handlers

import (
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/infrastructure/websocket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	Service          services.IDestinationService
	WebSocketManager *websocket.WebSocketManager
}

func (wc *WebSocketHandler) HandleConnections(c *gin.Context) {
	ws, err := websocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set WebSocket upgrade: " + err.Error()})
		return
	}

	defer func(ws *websocket.Connection) {
		err := ws.Close()
		if err != nil {
			return
		}

	}(ws)

	wc.WebSocketManager.AddWebSocketClient(ws)

	for {
		var msg websocket.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading json: %v", err)
			break
		}

		switch msg.Action {

		case "GenerateDestination":
			log.Printf("Got request to create destination")
			newDestination := msg.Destination
			//destination, err := wc.Service.CreateDestination(newDestination)
			//if err != nil {
			//	log.Printf("Error creating destination: %v", err)
			//	err := ws.WriteJSON(gin.H{"error": err.Error()})
			//	if err != nil {
			//		return
			//	}
			//	continue
			//}
			//wc.WebSocketManager.AddToBroadcast(websocket.EventUpdateNotification{Action: "CreateDestination", Destination: newDestination})
			err = ws.WriteJSON(newDestination)
			if err != nil {
				log.Printf("Error writing json: %v", err)
				return
			}
		}
	}

}
