package mocks

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"github.com/jaswdr/faker"
	"time"
)

type MockDestinationService struct {
	AllDestinationsFunc   func() ([]entities.Destination, error)
	DestinationByIDFunc   func(idStr string) (*entities.Destination, error)
	CreateDestinationFunc func(destination entities.Destination) (entities.Destination, error)
	UpdateDestinationFunc func(idStr string, updatedDestination entities.Destination) (entities.Destination, error)
	DeleteDestinationFunc func(idStr string) (entities.Destination, error)
}

func (m *MockDestinationService) DestinationsByLocationID(locationIDStr string) (*services.DestinationsByLocation, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDestinationService) StartGeneratingDestinations(interval time.Duration, f faker.Faker) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDestinationService) StopGeneratingDestinations() {
	//TODO implement me
	panic("implement me")
}

func (m *MockDestinationService) GenerateFakeDestination(f faker.Faker) (entities.Destination, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockDestinationService) AllDestinations() ([]entities.Destination, error) {
	return m.AllDestinationsFunc()
}

func (m *MockDestinationService) DestinationByID(idStr string) (*entities.Destination, error) {
	return m.DestinationByIDFunc(idStr)
}

func (m *MockDestinationService) CreateDestination(destination entities.Destination) (entities.Destination, error) {
	return m.CreateDestinationFunc(destination)
}

func (m *MockDestinationService) UpdateDestination(idStr string, updatedDestination entities.Destination) (entities.Destination, error) {
	return m.UpdateDestinationFunc(idStr, updatedDestination)
}

func (m *MockDestinationService) DeleteDestination(idStr string) (entities.Destination, error) {
	return m.DeleteDestinationFunc(idStr)
}
