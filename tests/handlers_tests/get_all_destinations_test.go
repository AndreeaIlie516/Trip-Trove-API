package handlers_test

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"Trip-Trove-API/tests/mocks"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAllDestinations_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	_ = &mocks.MockLocationService{
		LocationByIDFunc: func(idStr string) (*entities.Location, error) {
			return &entities.Location{
				Model:       gorm.Model{ID: 1},
				Name:        "Finnish Lapland",
				Country:     "Finland",
				Description: "Beautiful northern landscapes with aurora",
			}, nil
		},
	}

	mockService := &mocks.MockDestinationService{
		AllDestinationsFunc: func() ([]entities.Destination, error) {
			return []entities.Destination{
				{Name: "Beach Paradise", LocationID: 1},
			}, nil
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var destinations []entities.Destination
	err := json.Unmarshal(w.Body.Bytes(), &destinations)
	assert.NoError(t, err)
	assert.Len(t, destinations, 1)
	assert.Equal(t, "Beach Paradise", destinations[0].Name)
}

func TestAllDestinations_EmptyList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockDestinationService{
		AllDestinationsFunc: func() ([]entities.Destination, error) {
			return []entities.Destination{}, nil
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, "[]", w.Body.String(), "Expected an empty JSON array")
}

func TestAllDestinations_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockDestinationService{
		AllDestinationsFunc: func() ([]entities.Destination, error) {
			return nil, errors.New("internal server error")
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAllDestinations_LargeDataSet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	var largeDestinations []entities.Destination
	for i := 0; i < 1000; i++ {
		largeDestinations = append(largeDestinations, entities.Destination{Name: fmt.Sprintf("Destination %d", i), LocationID: 1})
	}

	_ = &mocks.MockLocationService{
		LocationByIDFunc: func(idStr string) (*entities.Location, error) {
			return &entities.Location{
				Model:       gorm.Model{ID: 1},
				Name:        "Finnish Lapland",
				Country:     "Finland",
				Description: "Beautiful northern landscapes with aurora",
			}, nil
		},
	}

	mockService := &mocks.MockDestinationService{
		AllDestinationsFunc: func() ([]entities.Destination, error) {
			return largeDestinations, nil
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var destinations []entities.Destination
	err := json.Unmarshal(w.Body.Bytes(), &destinations)
	assert.NoError(t, err)
	assert.Len(t, destinations, 1000, "Expected 1000 destinations in the response")
}
