package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"Trip-Trove-API/tests/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteLocation_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		DeleteLocationFunc: func(id string) (entities.Location, error) {
			return entities.Location{Name: "Deleted Location"}, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/locations/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteLocation_InvalidIDFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		DeleteLocationFunc: func(id string) (entities.Location, error) {
			return entities.Location{}, errors.New("invalid ID format")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/locations/abc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP status code 400 for invalid ID format")
}

func TestDeleteLocation_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		DeleteLocationFunc: func(id string) (entities.Location, error) {
			return entities.Location{}, errors.New("location not found")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/locations/9999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Expected HTTP status code 404 for location not found")
}
