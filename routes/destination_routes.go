package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterDestinationRoutes(router *gin.Engine, destinationHandler *handlers.DestinationHandler) {
	destinationGroup := router.Group("/destinations")
	{
		destinationGroup.GET("/", destinationHandler.AllDestinations)
		destinationGroup.GET("/:id", destinationHandler.DestinationByID)
		destinationGroup.GET("/location/:locationId", destinationHandler.DestinationsByLocationID)
		destinationGroup.POST("/", destinationHandler.CreateDestination)
		destinationGroup.PUT("/:id", destinationHandler.UpdateDestination)
		destinationGroup.DELETE("/:id", destinationHandler.DeleteDestination)
		destinationGroup.HEAD("/", destinationHandler.Head)
		destinationGroup.GET("/start-generating-destinations", destinationHandler.StartGeneratingDestinationsHandler)
		destinationGroup.GET("/stop-generating-destinations", destinationHandler.StopGeneratingDestinationsHandler)
	}
}
