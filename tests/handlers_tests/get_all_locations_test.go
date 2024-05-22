package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"Trip-Trove-API/tests/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllLocations_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		AllLocationsFunc: func() ([]entities.Location, error) {
			return []entities.Location{
				{Name: "Finnish Lapland", Country: "Finland", Description: "Beautiful northern landscapes with aurora"},
			}, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var locations []entities.Location
	err := json.Unmarshal(w.Body.Bytes(), &locations)
	assert.NoError(t, err)
	assert.Len(t, locations, 1)
	assert.Equal(t, "Finnish Lapland", locations[0].Name)
}

func TestAllLocations_EmptyList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		AllLocationsFunc: func() ([]entities.Location, error) {
			return []entities.Location{}, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String(), "Expected an empty JSON array")
}

func TestAllLocations_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		AllLocationsFunc: func() ([]entities.Location, error) {
			return nil, errors.New("internal server error")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAllLocations_LargeDataSet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	var largeLocations []entities.Location
	for i := 0; i < 1000; i++ {
		largeLocations = append(largeLocations, entities.Location{Name: "Location " + fmt.Sprintf("%d", i), Country: "Finland", Description: "Part of a large dataset"})
	}

	mockService := &mocks.MockLocationService{
		AllLocationsFunc: func() ([]entities.Location, error) {
			return largeLocations, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var locations []entities.Location
	err := json.Unmarshal(w.Body.Bytes(), &locations)
	assert.NoError(t, err)
	assert.Len(t, locations, 1000, "Expected 1000 locations in the response")
}
