package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoutes(router *gin.Engine, wsHandler *handlers.WebSocketHandler) {
	router.GET("/ws", wsHandler.HandleConnections)
}
