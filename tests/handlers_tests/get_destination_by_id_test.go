package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"Trip-Trove-API/tests/mocks"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDestinationByID_Success(t *testing.T) {
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
		DestinationByIDFunc: func(idStr string) (*entities.Destination, error) {
			return &entities.Destination{Name: "Beach Paradise", LocationID: 1}, nil
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var destination entities.Destination
	err := json.Unmarshal(w.Body.Bytes(), &destination)
	assert.NoError(t, err)
	assert.Equal(t, "Beach Paradise", destination.Name)
}

func TestDestinationByID_InvalidIDFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockDestinationService{
		DestinationByIDFunc: func(idStr string) (*entities.Destination, error) {
			return nil, errors.New("invalid ID format")
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/abc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDestinationByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockDestinationService{
		DestinationByIDFunc: func(idStr string) (*entities.Destination, error) {
			return nil, errors.New("destination not found")
		},
	}

	destinationHandler := &handlers.DestinationHandler{Service: mockService}
	routes.RegisterDestinationRoutes(router, destinationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/destinations/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
