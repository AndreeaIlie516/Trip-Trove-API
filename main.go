package main

import (
	"Trip-Trove-API/database"
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/infrastructure/dataaccess"
	"Trip-Trove-API/infrastructure/middlewares"
	"Trip-Trove-API/infrastructure/websocket"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()

	entitiesToMigrate := []interface{}{
		&entities.Destination{},
		&entities.Location{},
		&entities.User{},
	}

	for _, entity := range entitiesToMigrate {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	websocketManager := websocket.NewWebSocketManager()

	go websocketManager.BroadcastWebSocketMessage()

	authMiddleware := middlewares.AuthMiddleware{}

	destinationRepository := dataaccess.NewGormDestinationRepository(db)
	locationRepository := dataaccess.NewGormLocationRepository(db)
	userRepository := dataaccess.NewGormUserRepository(db)

	destinationService := services.DestinationService{Repo: destinationRepository, LocationRepo: locationRepository, WsManager: websocketManager}
	locationService := services.LocationService{Repo: locationRepository, DestinationRepo: destinationRepository}
	userService := services.UserService{Repo: userRepository}

	destinationHandler := handlers.DestinationHandler{Service: &destinationService}
	locationHandler := handlers.LocationHandler{Service: &locationService}
	userHandler := handlers.UserHandler{Service: &userService}

	routes.RegisterDestinationRoutes(router, &destinationHandler, authMiddleware)
	routes.RegisterLocationRoutes(router, &locationHandler, authMiddleware)
	routes.RegisterUserRoutes(router, &userHandler, authMiddleware)

	wsController := handlers.WebSocketHandler{Service: &destinationService, WebSocketManager: websocketManager}
	routes.RegisterWebSocketRoutes(router, &wsController)

	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
