package mocks

import "Trip-Trove-API/domain/entities"

type MockLocationService struct {
	AllLocationsFunc   func() ([]entities.Location, error)
	LocationByIDFunc   func(idStr string) (*entities.Location, error)
	CreateLocationFunc func(location entities.Location) (entities.Location, error)
	DeleteLocationFunc func(idStr string) (entities.Location, error)
	UpdateLocationFunc func(idStr string, location entities.Location) (entities.Location, error)
}

func (m *MockLocationService) AllLocations() ([]entities.Location, error) {
	return m.AllLocationsFunc()
}

func (m *MockLocationService) LocationByID(idStr string) (*entities.Location, error) {
	return m.LocationByIDFunc(idStr)
}

func (m *MockLocationService) CreateLocation(location entities.Location) (entities.Location, error) {
	return m.CreateLocationFunc(location)
}

func (m *MockLocationService) UpdateLocation(idStr string, updatedLocation entities.Location) (entities.Location, error) {
	return m.UpdateLocationFunc(idStr, updatedLocation)
}

func (m *MockLocationService) DeleteLocation(idStr string) (entities.Location, error) {
	return m.DeleteLocationFunc(idStr)
}
