package routes

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/infrastructure/middlewares"
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLocationRoutes(router *gin.Engine, locationHandler *handlers.LocationHandler, roleMiddleware middlewares.IAuthMiddleware) {
	locationGroup := router.Group("/locations")
	{
		locationGroup.GET("/", locationHandler.AllLocations)
		locationGroup.GET("/:id", locationHandler.LocationByID)
		locationGroup.POST("/", roleMiddleware.RequireRole(entities.Manager), locationHandler.CreateLocation)
		locationGroup.PUT("/:id", roleMiddleware.RequireRole(entities.Manager), locationHandler.UpdateLocation)
		locationGroup.DELETE("/:id", roleMiddleware.RequireRole(entities.Manager), locationHandler.DeleteLocation)
	}
}
