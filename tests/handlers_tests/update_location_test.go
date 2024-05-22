package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"Trip-Trove-API/tests/mocks"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateLocation_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		UpdateLocationFunc: func(id string, location entities.Location) (entities.Location, error) {
			return location, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	updatedLocation := entities.Location{
		Name:        "Updated Finnish Lapland",
		Country:     "Finland",
		Description: "Updated description of beautiful northern landscapes.",
	}
	requestBody, _ := json.Marshal(updatedLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/locations/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateLocation_NameTooShort(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		UpdateLocationFunc: func(id string, location entities.Location) (entities.Location, error) {
			return entities.Location{}, fmt.Errorf("Field 'name' validation failed")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	updatedLocation := entities.Location{
		Name:        "U",
		Country:     "Finland",
		Description: "Updated description of beautiful northern landscapes.",
	}
	requestBody, _ := json.Marshal(updatedLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/locations/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateLocation_InvalidCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		UpdateLocationFunc: func(id string, location entities.Location) (entities.Location, error) {
			return entities.Location{}, fmt.Errorf("Field 'country' validation failed")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	updatedLocation := entities.Location{
		Name:        "Updated Finnish Lapland",
		Country:     "",
		Description: "Updated description of beautiful northern landscapes.",
	}
	requestBody, _ := json.Marshal(updatedLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/locations/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
