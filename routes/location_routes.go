package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLocationRoutes(router *gin.Engine, locationHandler *handlers.LocationHandler) {
	locationGroup := router.Group("/locations")
	{
		locationGroup.GET("/", locationHandler.AllLocations)
		locationGroup.GET("/:id", locationHandler.LocationByID)
		locationGroup.POST("/", locationHandler.CreateLocation)
		locationGroup.PUT("/:id", locationHandler.UpdateLocation)
		locationGroup.DELETE("/:id", locationHandler.DeleteLocation)
	}
}
