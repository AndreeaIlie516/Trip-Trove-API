package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type LocationHandler struct {
	Service services.ILocationService
}

func (handler *LocationHandler) AllLocations(c *gin.Context) {
	locations, err := handler.Service.AllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch locations"})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func (handler *LocationHandler) LocationByID(c *gin.Context) {
	id := c.Param("id")
	location, err := handler.Service.LocationByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, location)
}

func (handler *LocationHandler) CreateLocation(c *gin.Context) {
	var newLocation entities.Location

	if err := c.BindJSON(&newLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.RegisterValidation("name", utils.NameValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("country", utils.CountryValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("description", utils.DescriptionValidator)
	if err != nil {
		return
	}

	err = validate.Struct(newLocation)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	location, err := handler.Service.CreateLocation(newLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create location"})
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (handler *LocationHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	location, err := handler.Service.DeleteLocation(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, location)
}

func (handler *LocationHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	var updatedLocation entities.Location

	if err := c.BindJSON(&updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.RegisterValidation("name", utils.NameValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("country", utils.CountryValidator)
	if err != nil {
		return
	}
	err = validate.RegisterValidation("description", utils.DescriptionValidator)
	if err != nil {
		return
	}

	err = validate.Struct(updatedLocation)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	location, err := handler.Service.UpdateLocation(id, updatedLocation)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}
