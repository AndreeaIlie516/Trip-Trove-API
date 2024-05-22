package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jaswdr/faker"
	"net/http"
	"strconv"
	"time"
)

type DestinationHandler struct {
	Service services.IDestinationService
}

func (handler *DestinationHandler) AllDestinations(c *gin.Context) {
	destinations, err := handler.Service.AllDestinations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch destinations"})
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func (handler *DestinationHandler) DestinationByID(c *gin.Context) {
	id := c.Param("id")
	destination, err := handler.Service.DestinationByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "destination not found"})
		}
		return
	}
	c.JSON(http.StatusOK, destination)
}

func (handler *DestinationHandler) DestinationsByLocationID(c *gin.Context) {
	locationID := c.Param("locationId")
	destinations, err := handler.Service.DestinationsByLocationID(locationID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func (handler *DestinationHandler) CreateDestination(c *gin.Context) {
	var newDestination entities.Destination

	if err := c.BindJSON(&newDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.RegisterValidation("name", utils.NameValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("description", utils.DescriptionValidator)
	if err != nil {
		return
	}

	if err := validate.Struct(newDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	destination, err := handler.Service.CreateDestination(newDestination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create destination"})
		return
	}

	//handler.Service.wsManager.AddToBroadcast(websocket.EventUpdateNotification{Action: "CreateDestination", Destination: destination})

	c.JSON(http.StatusCreated, destination)
}

func (handler *DestinationHandler) DeleteDestination(c *gin.Context) {
	id := c.Param("id")

	destination, err := handler.Service.DeleteDestination(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "destination not found"})
		}
		return
	}

	c.JSON(http.StatusOK, destination)
}

func (handler *DestinationHandler) UpdateDestination(c *gin.Context) {
	id := c.Param("id")

	var updatedDestination entities.Destination

	if err := c.BindJSON(&updatedDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.RegisterValidation("name", utils.NameValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("description", utils.DescriptionValidator)
	if err != nil {
		return
	}

	if err := validate.Struct(&updatedDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	destination, err := handler.Service.UpdateDestination(id, updatedDestination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update destination"})
		return
	}

	c.JSON(http.StatusOK, destination)
}

func (handler *DestinationHandler) Head(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.Status(http.StatusOK)
}

func (handler *DestinationHandler) StartGeneratingDestinationsHandler(c *gin.Context) {
	intervalParam := c.Query("interval")
	interval, err := strconv.Atoi(intervalParam)
	if err != nil || interval <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval parameter"})
		return
	}

	f := faker.New()
	handler.Service.StartGeneratingDestinations(time.Duration(interval)*time.Second, f)
	c.JSON(http.StatusOK, gin.H{"message": "Started generating destinations"})
}

func (handler *DestinationHandler) StopGeneratingDestinationsHandler(c *gin.Context) {
	println("Stop called")
	handler.Service.StopGeneratingDestinations()
	c.JSON(http.StatusOK, gin.H{"message": "Stopped generating destinations"})
}
