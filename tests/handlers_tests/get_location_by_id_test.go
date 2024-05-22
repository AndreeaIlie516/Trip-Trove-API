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

func TestLocationByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		LocationByIDFunc: func(idStr string) (*entities.Location, error) {
			return &entities.Location{
				Model:       gorm.Model{ID: 1},
				Name:        "Finnish Lapland",
				Country:     "Finland",
				Description: "Beautiful northern landscapes with aurora",
			}, nil
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var location entities.Location
	err := json.Unmarshal(w.Body.Bytes(), &location)
	assert.NoError(t, err)
	assert.Equal(t, "Finnish Lapland", location.Name)
}

func TestLocationByID_InvalidIDFormat(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		LocationByIDFunc: func(idStr string) (*entities.Location, error) {
			return nil, errors.New("invalid ID format")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/abc", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLocationByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &mocks.MockLocationService{
		LocationByIDFunc: func(idStr string) (*entities.Location, error) {
			return nil, errors.New("location not found")
		},
	}

	locationHandler := &handlers.LocationHandler{Service: mockService}
	routes.RegisterLocationRoutes(router, locationHandler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/locations/9999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
