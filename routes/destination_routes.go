package routes

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/infrastructure/middlewares"
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterDestinationRoutes(router *gin.Engine, destinationHandler *handlers.DestinationHandler, roleMiddleware middlewares.IAuthMiddleware) {
	destinationGroup := router.Group("/destinations")
	{
		destinationGroup.GET("/", destinationHandler.AllDestinations)
		destinationGroup.GET("/:id", destinationHandler.DestinationByID)
		destinationGroup.GET("/location/:locationId", destinationHandler.DestinationsByLocationID)
		destinationGroup.POST("/", roleMiddleware.RequireRole(entities.Manager), destinationHandler.CreateDestination)
		destinationGroup.PUT("/:id", roleMiddleware.RequireRole(entities.Manager), destinationHandler.UpdateDestination)
		destinationGroup.DELETE("/:id", roleMiddleware.RequireRole(entities.Manager), destinationHandler.DeleteDestination)
		destinationGroup.HEAD("/", destinationHandler.Head)
		destinationGroup.GET("/start-generating-destinations", destinationHandler.StartGeneratingDestinationsHandler)
		destinationGroup.GET("/stop-generating-destinations", destinationHandler.StopGeneratingDestinationsHandler)
	}
}
