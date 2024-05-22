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

func TestCreateLocation_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		CreateLocationFunc: func(location entities.Location) (entities.Location, error) {
			location.ID = 1
			return location, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	newLocation := entities.Location{
		Name:        "Finnish Lapland",
		Country:     "Finland",
		Description: "Beautiful northern landscapes with aurora",
	}
	requestBody, _ := json.Marshal(newLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/locations/", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var returnedLocation entities.Location
	err := json.Unmarshal(w.Body.Bytes(), &returnedLocation)
	assert.NoError(t, err)
	assert.Equal(t, newLocation.Name, returnedLocation.Name)
}

func TestCreateLocation_NameTooShort(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		CreateLocationFunc: func(location entities.Location) (entities.Location, error) {
			return entities.Location{}, fmt.Errorf("Field 'name' validation failed")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	newLocation := entities.Location{
		Name:        "Fi",
		Country:     "Finland",
		Description: "Beautiful northern landscapes with aurora",
	}
	requestBody, _ := json.Marshal(newLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/locations/", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateLocation_InvalidCountry(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		CreateLocationFunc: func(location entities.Location) (entities.Location, error) {
			return entities.Location{}, fmt.Errorf("Field 'country' validation failed")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	newLocation := entities.Location{
		Name:        "Finnish Lapland",
		Country:     "",
		Description: "Beautiful northern landscapes with aurora",
	}
	requestBody, _ := json.Marshal(newLocation)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/locations/", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
