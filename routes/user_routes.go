package routes

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/infrastructure/middlewares"
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler, roleMiddleware middlewares.IAuthMiddleware) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", roleMiddleware.RequireRole(entities.Admin), userHandler.AllUsers)
		userGroup.GET("/:id", roleMiddleware.RequireRole(entities.NormalUser), userHandler.UserByID)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
		userGroup.PUT("/:id", roleMiddleware.RequireRole(entities.NormalUser), userHandler.UpdateUser)
		userGroup.DELETE("/:id", roleMiddleware.RequireRole(entities.NormalUser), userHandler.DeleteUser)
	}
}
